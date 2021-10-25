package main

func searchMatrix(matrix [][]int, target int) bool {
	lengthx := len(matrix)
	lengthy := len(matrix[0])
	i, j := lengthx-1, lengthy-1
	for {
		if i < 0 {
			return false
		}
		if j >= lengthy {
			return false
		}
		if matrix[i][j] == target {
			return true
		}
		if target > matrix[i][j] {
			j++
			continue
		}
		if target < matrix[i][j] {
			i--
			continue
		}
	}
}
