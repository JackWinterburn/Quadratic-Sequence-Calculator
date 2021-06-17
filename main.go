package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
	"errors"
)

var BITSIZE int = 64

type NthTerm struct {
	OrigionalSequence     []float64
	HalfSecondDifference  float64
	Coefficient           float64
	Mutator               float64
}

func main() {
	nthTerm := NthTerm{}
	seq := getSequence()
	nthTerm.OrigionalSequence = getSequenceNums(seq)
	secondDifference, err := calcSteps(nthTerm.OrigionalSequence)
	if err != nil {
		log.Fatal(err)
	}

	nthTerm.HalfSecondDifference = float64(secondDifference) / 2

	fmt.Printf("\n Half second difference: %f", nthTerm.HalfSecondDifference)

	nthTerm.Coefficient, nthTerm.Mutator = getCoefficient(nthTerm.HalfSecondDifference, nthTerm.OrigionalSequence)

	fmt.Printf("\n The nth term is: %vn^2 + %vn + %v \n",
				nthTerm.HalfSecondDifference,nthTerm.Coefficient, nthTerm.Mutator)
}

func getSequence() string {
	fmt.Println("Please input a sequence (seperate the numbers with spaces)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sequence := scanner.Text()
	return sequence
}

// Turns the sequence string provided by the user, to an []float64array of floats,
func getSequenceNums(sequence string) []float64 {
	var numberSequence []float64
	splittedSequence := strings.Split(sequence, " ")

	for i := range splittedSequence {
		currNum, err := strconv.ParseFloat(splittedSequence[i], BITSIZE)
		if err != nil {
			log.Fatal("An error has occured. Please be sure to only input numbers with a single space in-between.")
		}
		numberSequence = append(numberSequence, currNum)
	}

	if len(numberSequence) < 4 { log.Fatal("We require 4 numbers to calculate the nth-term") }

	return numberSequence
}

func calcSteps(numSeq []float64) (float64, error) {
	var firstSteps []float64
	var secondSteps []float64

	// get the first difference
	for i := range numSeq {
		if i < len(numSeq) - 1 {
			difference := numSeq[i+1] - numSeq[i]
			firstSteps = append(firstSteps, difference) 
		}
	}

	// get the second difference
	for j := range firstSteps {
		if j < len(firstSteps) - 1 {
			difference := firstSteps[j+1] - firstSteps[j]
			secondSteps = append(secondSteps, difference)
		}
	}


	// compare second differences to ensure we are working with a quadratic sequence
	sequenceIsQuadratic := true
	if len(secondSteps) == 2 {
		if secondSteps[0] != secondSteps[1] {
			sequenceIsQuadratic = false
		}
	} else {
		for k := 0; k < len(secondSteps) - 1; k++ {
			if !sequenceIsQuadratic {
				break
			} else if secondSteps[k] != secondSteps[k+1]{
				sequenceIsQuadratic = false
			}
		}
	}

	if sequenceIsQuadratic {
		return secondSteps[0], nil
	} else {
		return 0, errors.New("Sequence provided was not quadratic.")
	}
}

func getCoefficient(hsd float64, ogSequence []float64) (float64, float64) {
	secondSequence := []float64{hsd * 1, hsd * 4, hsd * 9}

	fmt.Printf("\n Secondary (linear) sequence: %v", secondSequence)

	return calcLinearStep(secondSequence, ogSequence)
}

func calcLinearStep(seq []float64, ogs []float64) (float64, float64) {
	ogs = ogs[:3]
	diffBetOgAndSec := []float64{ogs[0] - seq[0], ogs[1] - seq[1], ogs[2] - seq[2]}
	fmt.Printf("\n Sequence of OG Sequence - Linear: %v", diffBetOgAndSec)

	firstStep := diffBetOgAndSec[1] - diffBetOgAndSec[0]
	secondStep := diffBetOgAndSec[2] - diffBetOgAndSec[1]

	if firstStep == secondStep {
		fmt.Printf("\n The linear sequence step is %f", firstStep)
		mutator := diffBetOgAndSec[0] - firstStep
		return firstStep, mutator
	} else {
		log.Fatal("Something went wrong whilst calculating the LINEAR sequence steps.")
	}
	return 0.0, 0.0
}