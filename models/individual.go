package models

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Action defines the possible alleles (actions) for each gene (state)
type Action int

const (
	MoveNorth Action = iota
	MoveSouth
	MoveEast
	MoveWest
	StayPut
	BendOver
	MoveRandom
)

// Actions lists all available actions to an individual
var actions = [...]Action{MoveNorth, MoveSouth, MoveEast, MoveWest, StayPut, BendOver, MoveRandom}

// Moves lists a subset of actions that are resolved move directions
var moves = [...]Action{MoveNorth, MoveSouth, MoveEast, MoveWest}

// Genome consists of actions assigned to each of the 243 possible states
type Genome [243]Action

// Individual defines one population member, their Genome and associated data
type Individual struct {
	Fitness float64
	Genome  *Genome
}

func (ind Individual) String() string {
	genome := ""
	for _, gene := range ind.Genome {
		genome += strconv.Itoa(int(gene))
	}
	return fmt.Sprintf("Fitness:\t%v\nGenome:\t%v", ind.Fitness, genome)
}

// RandomAction produces a single random action allele
func RandomAction() Action {
	return actions[rand.Intn(len(actions))]
}

// RandomMove produces a random resolved movement action
func RandomMove() Action {
	return moves[rand.Intn(len(moves))]
}

func randomGenome() *Genome {
	genome := &Genome{}
	for i := range genome {
		genome[i] = RandomAction()
	}
	return genome
}

// MakeIndividual creates a new Individual with a random Genome
func MakeIndividual() *Individual {
	return &Individual{
		Genome: randomGenome(),
	}
}
