package model

import (
	"math/rand"
	"time"
)

type BingoMatrix struct {
	cell [][]int
}

func initMatrix() BingoMatrix {

	var a []int
	var b BingoMatrix

	for i := 1; i <= 25; i++ {
		a = append(a, i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {

			b.cell[i][j] = a[(i*5)+j]
		}

	}

	return b
}
