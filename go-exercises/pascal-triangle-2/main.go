package main

import (
	"fmt"
)

func main() {

	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(21))
}

func getRow(rowIndex int) []int {
	// res := make([]int, rowIndex+1)
	//
	// rowIndex64 := uint64(rowIndex)
	// for i := uint64(0); i <= rowIndex64; i++ {
	// 	res[i] = int(calculateCombinatorialValue(rowIndex64, i))
	// }
	//
	// return res

	// res := make([][]int, rowIndex+1)
	//
	// res[0] = make([]int, 1)
	// res[0][0] = 1
	// for i := 1; i < rowIndex+1; i++ {
	//
	// 	res[i] = make([]int, i+1)
	// 	res[i][0] = 1
	// 	for j := 1; j < i; j++ {
	// 		res[i][j] = res[i-1][j-1] + res[i-1][j]
	// 	}
	// 	res[i][i] = 1
	// }
	// return res[rowIndex]

	res := make([]int, rowIndex+1)

	res[0] = 1

	for i := 1; i < rowIndex+1; i++ {
		for j := i; j > 0; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}

func calculateCombinatorialValue(n uint64, k uint64) uint64 {
	nMinusKPerm := n - k
	nMinusK := nMinusKPerm - 1

	nPerm := n
	n -= 1
	for n > 1 {
		nPerm *= n
		n -= 1
	}

	kPerm := k
	k -= 1
	for k > 1 {
		kPerm *= k
		k -= 1
	}

	for nMinusK > 1 {
		nMinusKPerm *= nMinusK
		nMinusK -= 1
	}

	if nPerm == 0 {
		nPerm = 1
	}

	if kPerm == 0 {
		kPerm = 1
	}

	if nMinusKPerm == 0 {
		nMinusKPerm = 1
	}

	return nPerm / (kPerm * nMinusKPerm)
}
