package main

import (
	"fmt"
	"math"
)

func normalizeMatrix(m [][]int) [][]float64 {
	normaliationConst := math.Pow(float64(len(m)), 2)

	var newMatrix [][]float64 = make([][]float64, len(m))

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			newMatrix[i] = append(newMatrix[i], float64(m[i][j])/(normaliationConst))
		}
	}
	return newMatrix
}

func PrintMatrix(m [][]float64) {
	fmt.Printf("%d\n", len(m))
	for i := 0; i < len(m); i++ {
		fmt.Print("[")
		for j := 0; j < len(m); j++ {
			fmt.Printf("%.4f ", m[i][j])
		}
		fmt.Print("]\n")
	}
}

func generateNewMatrix(previous [][]int) [][]int {

	var previousSize int = len(previous)
	var newMatrix [][]int = make([][]int, previousSize*2)
	var lineHolder1 [][]int = make([][]int, previousSize)
	var lineHolder2 [][]int = make([][]int, previousSize)

	for i := 0; i < previousSize; i++ {
		for j := 0; j < previousSize; j++ {
			lineHolder1[i] = append(lineHolder1[i], (4*previous[i][j])+0)
			lineHolder2[i] = append(lineHolder2[i], (4*previous[i][j])+2)
		}
	}

	for i := 0; i < previousSize; i++ {
		for j := 0; j < previousSize; j++ {
			newMatrix[i] = append(newMatrix[i], lineHolder1[i][j])
		}
		for j := 0; j < previousSize; j++ {
			newMatrix[i] = append(newMatrix[i], lineHolder2[i][j])
		}
	}

	lineHolder1 = make([][]int, previousSize)
	lineHolder2 = make([][]int, previousSize)

	for i := 0; i < previousSize; i++ {
		for j := 0; j < previousSize; j++ {
			lineHolder1[i] = append(lineHolder1[i], (4*previous[i][j])+3)
			lineHolder2[i] = append(lineHolder2[i], (4*previous[i][j])+1)
		}
	}

	for i := 0; i < previousSize; i++ {
		for j := 0; j < previousSize; j++ {
			newMatrix[i+previousSize] = append(newMatrix[i+previousSize], lineHolder1[i][j])
		}
		for j := 0; j < previousSize; j++ {
			newMatrix[i+previousSize] = append(newMatrix[i+previousSize], lineHolder2[i][j])
		}
	}
	return newMatrix
}

func recursiveGenerateMatrix(n int) [][]int {

	if n == 0 {
		return [][]int{{0, 2}, {3, 1}}
	}

	var previous [][]int = recursiveGenerateMatrix(n - 1)
	return generateNewMatrix(previous)
}

func NewNormalizedMatrix(n int) [][]float64 {
	var bayerMatrix = recursiveGenerateMatrix(n)
	return normalizeMatrix(bayerMatrix)
}

func NewNonNormalizedMatrix(n int) [][]int {
	return recursiveGenerateMatrix(n)
}
