package main

import (
	"fmt"
	"math/rand"
)

// The official alphabet.
var ALPHABET = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ")


// GetFitness calculates the "fitness" of a string with respect to another,
// how similar they are in a character to character basis. Works with slice of runes
// directly to avoid multiple conversions. May either fail because rune slices
// are not the same size, or "fail" because of a clone.
func GetFitness(rTarget, rSon []rune) (int, error){
	if len(rTarget) != len(rSon) {
		return 0, fmt.Errorf("getting fitness: both rune slices must be of the same size");
	}

	fitness := 0

	for i := range rTarget {
		if rTarget[i] == rSon[i] {
			fitness++
		}
	}

	return fitness, nil
}

// Mutate derives a string from another one. The mutation rate is a percentage
// expressed in the range [0.0, 1.0]; it determines what is the probability
// of any given letter to "mutate". In theory, this should (shittily) simulate
// how DNA is copied with the ocassional random error Works directly with
// rune strings to avoid multiple conversions.
func Mutate(parentRunes []rune, mutationRate float64) []rune {
	result := make([]rune, len(parentRunes))

	for i := range parentRunes {
		if rand.Float64() < mutationRate {
			result[i] = ALPHABET[ rand.Intn(len(ALPHABET)) ]
		} else {
			result[i] = parentRunes[i]
		}
	}

	return result
}

// Completely random rune slice formed by the offcial alphabet
func RandomRuneSlice(length int) []rune {
	result := make([]rune, length)

	for i := range result {
		result[i] = ALPHABET[ rand.Intn(len(ALPHABET)) ]
	}

	return result
}

// FindFittest finds the fittest Offspring in an OffspringSlice. Returns
// both the index and a pointer to the Offspring itself.
//
// It's a linear search, so in case of multiple fittest Offsprings, it will
// return the first found, but will have searched the entire OffspringSlice to make
// sure. In case of an Overman, it will return immediately.
func FindFittest(offs OffspringSlice) (int, *Offspring) {
	maxFitnessFound := 0
	fittest := 0
	for i, currOff := range offs {
		if currOff.Overman() {
			return i, currOff
		}

		if currOff.Fitness > maxFitnessFound {
			maxFitnessFound = currOff.Fitness
			fittest = i;
		}
	}

	return fittest, offs[fittest];
}
