package fitness

import (
	"robbie/models"
)

const trials = 50

const (
	currentBase = 1
	northBase   = 3
	southBase   = 9
	eastBase    = 27
	westBase    = 81
)

func getSensedState(floorMap *models.FloorMap, x int, y int) int {
	return int(northBase*floorMap[y-1][x] +
		southBase*floorMap[y+1][x] +
		eastBase*floorMap[y][x+1] +
		westBase*floorMap[y][x-1] +
		currentBase*floorMap[y][x])
}

var trainMaps = makeTrainingMaps()

func makeTrainingMaps() []models.FloorMap {
	var trainMaps []models.FloorMap
	for i := 0; i < trials; i++ {
		trainMaps = append(trainMaps, models.MakeFloorMap())
	}
	return trainMaps
}

// CheckFitness gets one fitness of an Individual on a FloorMap
func checkFitness(ind *models.Individual, floorMap models.FloorMap) int {
	x, y := 1, 1
	fitness := 0

	for step := 0; step < 200; step++ {
		state := getSensedState(&floorMap, x, y)
		action := ind.Genome[state]

		if action == models.StayPut {
			continue
		} else if action == models.BendOver {
			if floorMap[y][x] == models.Can {
				floorMap[y][x] = models.Empty
				fitness += 10
			} else {
				fitness--
			}
			continue
		} else if action == models.MoveRandom {
			action = models.RandomMove()
		}

		if action == models.MoveNorth && floorMap[y-1][x] != models.Wall {
			y--
		} else if action == models.MoveSouth && floorMap[y+1][x] != models.Wall {
			y++
		} else if action == models.MoveEast && floorMap[y][x+1] != models.Wall {
			x++
		} else if action == models.MoveWest && floorMap[y][x-1] != models.Wall {
			x--
		} else {
			// Ran into a wall
			fitness -= 5
		}
	}

	return fitness
}

// avgFitness assigns individual Fitness mean based on multiple trials
func avgFitness(ind *models.Individual) {
	totalFitness := 0

	for i := 0; i < trials; i++ {
		// trainMap := trainMaps[i]
		// totalFitness += checkFitness(ind, trainMap)
		totalFitness += checkFitness(ind, models.MakeFloorMap())
	}

	ind.Fitness = float64(totalFitness) / float64(trials)
}

// PopFitness gets the average fitness of an entire population
func PopFitness(pop *models.Population) float64 {
	totalFitness := 0.0

	for i := range pop {
		avgFitness(pop[i])
		totalFitness += pop[i].Fitness
	}

	return totalFitness / float64(len(pop))
}
