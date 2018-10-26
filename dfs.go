package main

import (
	"math"
)

var rowNum = []int{0, 1, 1}
var colNum = []int{1, 1, 0}

func isMutant(dna [6]string) bool {
	for i, s := range dna {
		for j, currentNode := range s {
			var founded = dfsRecur(dna, i, j, -1, -1, currentNode, 0)

			if founded {
				return true
			}
		}
	}

	return false
}

func dfsRecur(matrix [6]string, row int, col int, prevRow int, prevCol int, node rune, counter int) bool {
	var initialNode = string(node)
	var currentNode = string(matrix[row][col])

	if currentNode != initialNode {
		return false
	}

	if counter == 3 {
		return true
	}

	for k := 0; k < 3; k++ {
		if isValidMovement(row, col, row+rowNum[k], col+colNum[k], prevRow, prevCol, counter) {
			var founded = dfsRecur(matrix, row+rowNum[k], col+colNum[k], row, col, node, counter+1)

			if founded {
				return true
			}
		}
	}

	return false
}

func isValidMovement(currentRow int, currentColumn int, row int, col int, prevRow int, prevCol int, counter int) bool {
	if counter >= 1 {
		var distRow = math.Abs(float64(currentRow - prevRow))
		var distCol = math.Abs(float64(currentColumn - prevCol))

		return (row >= 0) && (row < 6) && (col >= 0) && (col < 6) &&
			(row == currentRow+int(distRow)) && (col == currentColumn+int(distCol))
	}

	return (row >= 0) && (row < 6) && (col >= 0) && (col < 6) &&
		!(row == prevRow && col == prevCol)
}
