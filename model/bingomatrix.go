package model

import (
	"math/rand"
	"time"
)
//BingoMatrix Structure
type BingoMatrix struct {
	Cell [5][5]int  `json:"cell"`
}

//InitMatrix initializes a new matrix of a player
func (b *BingoMatrix) InitMatrix() {

	var a []int

	for i := 1; i <= 25; i++ {
		a = append(a, i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {

			b.Cell[i][j] = a[(i*5)+j]
		}

	}
}

//UpdateMatrix updates changes in the matrix
func (b *BingoMatrix) UpdateMatrix(val int) (int, int) {

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {

			if b.Cell[i][j] == val {
				b.Cell[i][j] = 0
				return i, j
			}
		}
	}
	return -1, -1

}
