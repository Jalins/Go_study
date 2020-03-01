package main

import "fmt"

type nodeVal struct {
	row int
	col int
	val int
}

func main() {
	var chessArray [11][11]int

	chessArray[1][2] = 1
	chessArray[2][3] = 2

	for _, value1 := range chessArray {
		for _, value2 := range value1 {
			fmt.Printf("%d\t", value2)
		}

		fmt.Println()
	}

	// 压缩数组
	var sparseArray []nodeVal
	var valNode nodeVal
	for i, value1 := range chessArray {
		for j, value2 := range value1 {
			if value2 != 0 {
				valNode = nodeVal{
					row: i,
					col: j,
					val: value2,
				}

				sparseArray = append(sparseArray, valNode)
			}

		}
	}

	for key, value := range sparseArray {
		fmt.Printf("%d: %d  %d  %d\n", key, value.row, value.col, value.val)
	}

	// 解压数组
	var newChessArray [11][11]int
	for _, value := range sparseArray {
		newChessArray[value.row][value.col] = value.val
	}

	for _, value := range newChessArray {
		for _, value1 := range value {
			fmt.Printf("%d\t", value1)
		}
		fmt.Println()
	}

}
