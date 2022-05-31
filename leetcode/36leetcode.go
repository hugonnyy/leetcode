package leetcode

// 请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
//
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//
//
//注意：
//
//一个有效的数独（部分已被填充）不一定是可解的。
//只需要根据以上规则，验证已经填入的数字是否有效即可。
//空白格用 '.' 表示。
//func isValidSudoku(board [][]byte) bool {
//	// 数组操作的速度比map快
//
//	var row, column [][]int
//	var rectangle [][][]int
//
//	for i := 0; i < 9; i++ {
//		for j := 0; j < 9; j++ {
//
//		}
//	}
//
//	row := make([]map[byte]bool, 9)
//	column := make([]map[byte]bool, 9)
//	rectangle := make([]map[byte]bool, 9)
//
//	for i := 0; i < 9; i++ {		// row
//		for j := 0; j < 9; j++ {	// column
//
//			if board[i][j] == '.' {
//				continue
//			}
//
//			rowMap := row[i]
//			if rowMap == nil {
//				rowMap = make(map[byte]bool)
//				row[i] = rowMap
//			}
//
//			if _, exist := rowMap[board[i][j]]; exist {
//				return false
//			}else {
//				rowMap[board[i][j]] = true
//			}
//
//			columnMap := column[j]
//			if columnMap == nil {
//				columnMap = make(map[byte]bool)
//				column[j] = columnMap
//			}
//
//			if _, exist := columnMap[board[i][j]]; exist {
//				return false
//			}else {
//				columnMap[board[i][j]] = true
//			}
//
//			index := 0
//			index += j / 3
//			index += (i / 3) * 3
//			rectangleMap := rectangle[index]
//			if rectangleMap == nil {
//				rectangleMap = make(map[byte]bool)
//				rectangle[index]= rectangleMap
//			}
//
//			if _, exist := rectangleMap[board[i][j]]; exist {
//				return false
//			}else {
//				rectangleMap[board[i][j]] = true
//			}
//		}
//	}
//
//	return true
//}

func isValidSudoku(board [][]byte) bool {
	// 数组操作的速度比map快
	// 类似与map的操作，将具体数值存到索引中，根据索引进行数据查找
	var rows, columns [9][9]int
	var rectangles [3][3][9]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				continue
			}

			index := b - '1'

			rows[i][index]++
			columns[j][index]++
			rectangles[i/3][j/3][index]++

			if rows[i][index] > 1 || columns[j][index] > 1 || rectangles[i/3][j/3][index] > 1 {
				return false
			}
		}
	}

	return true
}
