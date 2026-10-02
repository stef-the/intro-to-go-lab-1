package main

func findLiveNeighbors(p golParams, world [][]byte, y int, x int) int {
	// find all 8 adjacent for a cell
	offsetValues := []int{-1, 0, 1}
	livingCells := 0
	for _, yOffset := range offsetValues {
		for _, xOffset := range offsetValues {
			if yOffset == 0 && xOffset == 0 {
				continue
			}
			yCoordinate := (y + yOffset + p.imageHeight) % p.imageHeight
			xCoordinate := (x + xOffset + p.imageWidth) % p.imageWidth
			if world[yCoordinate][xCoordinate] == 255 {
				livingCells += 1
			}
		}
	}
	return livingCells
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	// create new world
	newWorld := make([][]byte, p.imageHeight)
	for y := range newWorld {
		newWorld[y] = make([]byte, p.imageWidth)
		for x := range newWorld[y] {
			currentLivingNeighbors := findLiveNeighbors(p, world, y, x)
			if world[y][x] == 0 && currentLivingNeighbors == 3 {
				newWorld[y][x] = 255
			} else if world[y][x] != 0 && (currentLivingNeighbors == 2 || currentLivingNeighbors == 3) {
				newWorld[y][x] = 255
			} else {
				newWorld[y][x] = 0
			}
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var aliveCells []cell
	for y := range world {
		for x := range world[y] {
			if world[y][x] == 255 {
				aliveCells = append(aliveCells, cell{x: x, y: y})
			}
		}
	}
	return aliveCells
}
