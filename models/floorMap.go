package models

import (
	"fmt"
	"math/rand"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type FloorMap [12][12]Cell

func (floorMap FloorMap) String() string {
	visual := ""
	for _, row := range floorMap {
		visual += fmt.Sprintf("%v\n", row)
	}
	return visual
}

func MakeFloorMap() FloorMap {
	var floorMap FloorMap
	for yi, floorRow := range floorMap {
		for xi := range floorRow {
			if (xi == 0) || (yi == 0) || (xi == len(floorRow)-1) || (yi == len(floorMap)-1) {
				floorMap[yi][xi] = Wall
			} else if rand.Float64() < 0.5 {
				floorMap[yi][xi] = Can
			}
		}
	}
	return floorMap
}
