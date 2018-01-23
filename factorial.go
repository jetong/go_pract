// time test factorial on matrix of values

package main

import (
	"fmt"
	"time"
)

func fact(n int64, t time.Time) int64 {
	if n == 1 {
		fmt.Println(time.Since(t))
		return int64(1)
	} else {
		time.Sleep(1 * time.Second)
		return fact(n-1, t) * n
	}
}

// take factorial of each value in the matrix
func matrixFact(mat [][]int64, t time.Time) [][]int64 {
	length := len(mat)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			mat[i][j] = fact(mat[i][j], t)
		}
	}
	return mat
}

func main() {
	fmt.Print("Enter n for slice size: ")
	var value int64
	fmt.Scanf("%d", &value)
	fmt.Println("You entered:", value)

	// generate matrix
	matrix := make([][]int64, value)
	for i := range matrix {
		matrix[i] = make([]int64, value)
	}
	// fill with values from 1 to n
	count := int64(1)
	for i := 0; i < int(value); i++ {
		for j := 0; j < int(value); j++ {
			matrix[i][j] = count
			count++
		}
	}

	fmt.Println("Matrix before:", matrix)
	start := time.Now()
	fmt.Println("Timer start: ", time.Since(start))
	matrix = matrixFact(matrix, start)
	fmt.Println("Elapsed time:", time.Since(start))
	fmt.Println("Matrix after:", matrix)
}
