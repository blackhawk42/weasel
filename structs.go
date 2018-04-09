package main

import (
	"fmt"
)

// Offspring represents an offspring. In practice, this can be though of as an "organism",
// as it isn't *necessarily* derived from another.
type Offspring struct {
	// Parent of this offspring. May be nil if created from the dust on the ground. Yes,
	// this makes this structure a sort of weird "reverse" linked list.
	Parent *Offspring
	
	// The phrase of this offspring. Normally derived from the parent. It's a rune slice for
	// avoidance of unnecessary conversions and possible language extension.
	Phrase []rune
	
	// Target phrase, the platonic ideal we all aspire to be (to not die). Can be
	// used to get "max fitness", as that's simply the length of the phrase.
	TargetPhrase []rune
	
	// Generation this offspring belongs to. 0 is a convenient number for an Adam,
	// but hey, you do you.
	Generation int
	
	// Fitness in relationship to the parent. How many characters of the phrase
	// are the same with respect to the target phrase?
	Fitness int
}

// Create a first Offspring, who shall rule over the bytes on the memory and
// the calls on the OS.
func CreateAdam(targetPhrase []rune) (*Offspring, error) {
	phrase := RandomRuneSlice( len(targetPhrase) )
	fitness, err := GetFitness(targetPhrase, phrase)
	if err != nil {
		return nil, err
	}
	
	return &Offspring {
		Parent: nil,
		Phrase: phrase,
		TargetPhrase: targetPhrase,
		Generation: 0,
		Fitness: fitness,
	}, nil
}

// Parthogenesis a new offspring, taking mutation into account.
func (parent *Offspring) Spawn(mutationRate float64) (*Offspring, error) {
	son := &Offspring {
		Parent: parent,
		Phrase: Mutate(parent.Phrase, mutationRate),
		TargetPhrase: parent.TargetPhrase,
		Generation: parent.Generation + 1,
	}
	
	var err error
	son.Fitness, err = GetFitness(son.TargetPhrase, son.Phrase)
	if err != nil {
		return nil, err
	}
	
	return son, nil
}

// Spawn with concurrency
func (parent *Offspring) goSpawn(mutationRate float64, channel chan<- *BirthReport) {
	report := &BirthReport{}
	report.Newborn, report.Error = parent.Spawn(mutationRate)
	
	channel <- report
}

// Get max fitness possible for this organism
func (of *Offspring) MaxFitness() int {
	return len(of.TargetPhrase)
}

// Calculate a ratio between this offspring's fitness and maximum fitness, mostly
// for reporting purposes.
func (of *Offspring) RelativeFitness() float64 {
	return float64(of.Fitness) / float64( of.MaxFitness() )
}

// Detect if offspring has trascended morality and became a life-affirming Overman.
func (of *Offspring) Overman() bool {
	return of.Fitness == of.MaxFitness()
}

// Create a nice-looking report for printing
func (of *Offspring) Report() string {
	return fmt.Sprintf("%04d: %s  --- score: %4.2f%% (%d/%d)", of.Generation,
	string(of.Phrase), of.RelativeFitness()*100, of.Fitness, of.MaxFitness() )
}



// Slice of Offsprings. Implements Golang Sort Interface.
type OffspringSlice []*Offspring

// Convenience function to create an OffspringSlice
func NewOffspringSlice(length int) OffspringSlice {
	return OffspringSlice( make([]*Offspring, length) )
}

// Length of the slice.
func (s OffspringSlice) Len() int {
	return len(s)
}

// Tell if an offspring is "less" than other. Based on fitness.
func (s OffspringSlice) Less(i, j int) bool {
	return s[i].Fitness < s[j].Fitness
}

// Swap an offpring from another in the slice.
func (s OffspringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


// Represents a birth report, with a born offspring and any... miscarriages
// that could have happened. Mostly for use with concurrency.
type BirthReport struct {
	Newborn *Offspring
	Error error
}
