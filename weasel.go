package main

import (
	"fmt"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"math/rand"
	"time"
)

const DEFAULT_OFFSPRING_NUMBER int = 100
const DEFAULT_MUTATION_RATE float64 = 0.05
const DEFAULT_TARGET_FRASE string = "METHINKS IT IS LIKE A WEASEL"
const DEFAULT_MAX_GENERATIONS int = 1000

func main() {
	// Options and non-standard exit points
	var offspringNumber = flag.Int("offspring-per-generation", DEFAULT_OFFSPRING_NUMBER, "number of `offspring` generated per each generation")
	var mutationRate = flag.Float64("mutation-rate", DEFAULT_MUTATION_RATE, "mutation `rate`, expressed as a percentage [0.0, 1.0]")
	var targetPhrase = flag.String("target", DEFAULT_TARGET_FRASE, "target `phrase` aginst which fitness will be measured")
	var maxGenerations = flag.Int("max-generations", DEFAULT_MAX_GENERATIONS, "maximum number of `generations`, to avoid an endless orgy")
	
	flag.Usage = func () {
		fmt.Fprintf(os.Stderr, "use: %s [OPTIONS]\n\n", filepath.Base(os.Args[0]) )
		fmt.Printf("Slightly shitty version of Richard Dawkins' \"weasel\" algortihm, from \"The Blind Watchmaker\".\n\n")
		flag.PrintDefaults()
	}
	
	flag.Parse()
	
	
	if *offspringNumber < 1 {
		fmt.Fprintf(os.Stderr, "please use a positive integer for number of offsprings\n")
		flag.Usage()
		os.Exit(2)
	}
	
	if *mutationRate < 0.0 || *mutationRate > 1.0 {
		fmt.Fprintf(os.Stderr, "please use a mutation rate in the range [0.0, 1.0]\n")
		flag.Usage()
		os.Exit(2)
	}
	
	// Random
	
	rand.Seed(time.Now().UnixNano())
	
	// Mian logic
	
	
	fmt.Printf("Mutation rate: %.2f\n", *mutationRate)
	fmt.Printf("Number of offsprings per generation: %d\n", *offspringNumber)
	fmt.Printf("Max number of generations: %d\n", *maxGenerations)
	fmt.Printf("Target phrase: %s\n", *targetPhrase)
	fmt.Printf("\n")
	
	
	offsprings := NewOffspringSlice(*offspringNumber)
	fittest := len(offsprings) - 1
	
	var err error
	offsprings[ fittest ], err = CreateAdam( []rune(*targetPhrase) )
	if err != nil {
		fmt.Fprintf(os.Stderr, "creating Adam: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Adam: %s\n", offsprings[ fittest ].Report())
	
	if offsprings[ fittest ].Overman() {
		fmt.Printf("A miracle!!!\n")
		os.Exit(0)
	}
	
	
	for currentGeneration := 1; currentGeneration <= *maxGenerations; currentGeneration++ {
		parent := offsprings[ fittest ]
		
		for i := range offsprings {
			offsprings[i], err = parent.Spawn(*mutationRate)
			if err != nil {
				fmt.Fprintf(os.Stderr, "while spawinig children: %v\n", err)
				os.Exit(1)
			}
		}
		
		sort.Sort(offsprings)
		
		fmt.Printf("%s\n", offsprings[ fittest ].Report() )
		
		if offsprings[ fittest ].Overman() {
			break
		}
	}
	
	fmt.Printf("\n")
	
	if offsprings[ fittest ].Overman() {
		fmt.Printf("Overman found!\n")
	} else {
		fmt.Printf("Max number of generations reached\n")
	}
}
