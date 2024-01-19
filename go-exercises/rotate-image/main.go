package main

import "fmt"

func main() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9}
    matrix := [][]int{first, second, third}

    rotate(matrix)
    fmt.Println(matrix)
}

func rotate(matrix [][]int) {
    n, m := len(matrix), len(matrix[0])

    for i := 0; i < n; i++ {

        // transpose of row
        for j := i; j < m; j++ {
            matrix[i][j],matrix[j][i] = matrix[j][i], matrix[i][j]
        }

        // flipping the row
        for j, k := 0, m-1; j < k; j, k = j+1, k-1 {
            matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
        }
    }
}
