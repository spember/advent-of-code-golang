package aoc2024

import (
	"fmt"
	"slices"

	"github.com/spember/advent-of-code-golang/pkg/aocutils"
)

func NewGardenRegion(plant string) *GardenRegion {
	return &GardenRegion{
		Plant: plant,
		plots: make([][2]int, 0),
	}
}

type GardenRegion struct {
	Plant string
	plots [][2]int
}

func (g *GardenRegion) AddPlot(plot [2]int) {
	if !slices.Contains(g.plots, plot) {
		g.plots = append(g.plots, plot)
	}
}

func (g *GardenRegion) Area() int {
	return len(g.plots)
}

func (g *GardenRegion) Perimeter(garden [][]string) int {
	var perimeter int
	for _, plot := range g.plots {
		// for every plot, check each orthogonal direction. if different than Plant or off grid, increment perimeter
		// up
		dirs := [][2]int{
			{plot[0] - 1, plot[1]},
			{plot[0] + 1, plot[1]},
			{plot[0], plot[1] - 1},
			{plot[0], plot[1] + 1},
		}
		for _, dir := range dirs {
			if !aocutils.IsOnGrid(garden, dir[0], dir[1]) || garden[dir[0]][dir[1]] != g.Plant {
				perimeter++
			}
		}
	}
	return perimeter
}

func Day12Part1(input []byte) int {
	garden := aocutils.ParseToStringGrid(input)
	// scan for regions by tracking where we've visited, and then scanning around for the same symbol

	regions := make([]*GardenRegion, 0)
	visited := make(map[[2]int]bool)
	var currentRegion *GardenRegion

	for r, row := range garden {
		for c := 0; c < len(row); c++ {
			// if we've visited this plot, skip
			if visited[[2]int{r, c}] {
				continue
			}
			fmt.Println("Starting new region ", garden[r][c], " at ", r, c)
			currentRegion = NewGardenRegion(garden[r][c])
			//currentRegion.AddPlot([2]int{r, c})
			// search around the current plot
			plotsToCheck := [][2]int{{r, c}}

			for len(plotsToCheck) > 0 {
				//fmt.Println("starting plots to check", plotsToCheck)
				currentPlot := plotsToCheck[0]
				plotsToCheck = plotsToCheck[1:]

				// if plot is not in the garden, skip
				if currentPlot[0] < 0 || currentPlot[0] >= len(garden) || currentPlot[1] < 0 || currentPlot[1] >= len(garden[0]) {
					continue
				}

				// if plot is not the same as the current region, skip
				if garden[currentPlot[0]][currentPlot[1]] != currentRegion.Plant {
					continue
				}

				// if plot is the same as the current region, add it to the region
				currentRegion.AddPlot(currentPlot)
				visited[currentPlot] = true

				// add more plots to check
				for _, nextStep := range aocutils.FindNextOrthogonalSteps(garden, currentPlot[0], currentPlot[1]) {
					if !visited[nextStep] {
						plotsToCheck = append(plotsToCheck, nextStep)
					}
				}

			}
			fmt.Println("region", currentRegion.Plant, "has", currentRegion.Area(), "plots")
			regions = append(regions, currentRegion)
		}
	}

	var price = 0
	for _, region := range regions {
		fmt.Println(region.Plant, region.Area(), region.Perimeter(garden))
		price += region.Area() * region.Perimeter(garden)
	}

	return price
}

func Day12Part2(input []byte) int {
	// same as part 1, but we need to check for "sides" of each Garden region. Grow check contiguous permimeter.
	// No time today to do it.`
	return 0
}
