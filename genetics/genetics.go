package genetics

import (
	"math/rand"
	"robbie/models"
	"sort"
)

// NewGeneration creates a new population from the previous one
func NewGeneration(oldPop *models.Population) *models.Population {
	parents := selection(oldPop)
	children := reproduce(parents)
	for i := range children {
		mutate(children[i])
	}

	return children
}

func choices(population *models.Population, n int) []*models.Individual {
	perm := rand.Perm(len(population))
	var buffer []*models.Individual
	for i := 0; i < n; i++ {
		buffer = append(buffer, population[perm[i]])
	}
	return buffer
}

// byFitness sorts a group of Individuals by their Fitness
type byFitness []*models.Individual

func (a byFitness) Len() int           { return len(a) }
func (a byFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFitness) Less(i, j int) bool { return a[i].Fitness > a[j].Fitness }

// selection returns an entire population of parents
func selection(pop *models.Population) *models.Population {
	var selected models.Population

	for i := 0; i < len(pop); i += 2 {
		tourney := choices(pop, 11)
		sort.Slice(tourney, func(i, j int) bool {
			return tourney[i].Fitness > tourney[j].Fitness
		})
		// ranked := sort.Sort(byFitness(tourney))
		selected[i] = tourney[0]
		selected[i+1] = tourney[1]
	}

	return &selected
}

// reproduce creates a children population via crossover of a parent population
func reproduce(pop *models.Population) *models.Population {
	var children models.Population

	for i := 0; i < len(pop); i += 2 {
		parentA := pop[i]
		parentB := pop[i+1]

		split := rand.Intn(243)

		var a models.Genome
		var b models.Genome

		// Copy parent genomes into children by split point
		copy(a[:split], parentA.Genome[:split])
		copy(a[split:], parentB.Genome[split:])
		copy(b[:split], parentB.Genome[:split])
		copy(b[split:], parentA.Genome[split:])

		children[i] = &models.Individual{Genome: &a}
		children[i+1] = &models.Individual{Genome: &b}
	}

	return &children
}

func mutate(ind *models.Individual) {
	for i := range ind.Genome {
		if rand.Float64() < 0.02 {
			ind.Genome[i] = models.RandomAction()
		}
	}
}
