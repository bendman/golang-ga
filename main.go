package main

import (
	"fmt"
	"math/rand"
	"robbie/fitness"
	"robbie/genetics"
	"robbie/models"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var pop models.Population

	for i := range pop {
		pop[i] = models.MakeIndividual()
	}

	var best models.Individual
	oldGen := pop
	for i := 0; i < 2500; i++ {
		popFitness := fitness.PopFitness(&oldGen)
		sort.Slice(oldGen[:], func(i, j int) bool {
			return oldGen[i].Fitness > oldGen[j].Fitness
		})
		maxFitness := oldGen[0].Fitness
		fmt.Printf("\nGen %v\t\tMean: %v\tBest: %v", i, int(popFitness), maxFitness)
		if maxFitness > best.Fitness {
			best = *oldGen[0]
			fmt.Printf("\nNew Best: %v", best)
		}

		newGen := genetics.NewGeneration(&oldGen)
		oldGen = *newGen
	}

	fmt.Println(best)
}

// 054052254152050251356054251251250252251321255054561514156335152154602153356222305355352625355251605353351553224205035151001101353354404351114443354203444336551420255053251153616255053052151252225256310600254366620360304461634065612120612544653
