func isValidSudoku(board [][]byte) bool {
    var row [9][9]bool
    var col [9][9]bool
    var boxes [9][9]bool
    for i:=0; i<len(board); i++ {
        for j:=0; j<len(board[i]); j++ {
            if board[i][j] == '.' {
                continue
            }
            
            value := int(board[i][j]- '0')
            if value > 9 || value < 1 {
                return false
            }
            xBoxes := 3 * (i / 3) + (value - 1) / 3
            yBoxes := 3 * (j / 3) + (value - 1) % 3

            if row[i][value - 1] || col[value - 1][j] || boxes[xBoxes][yBoxes] {
                return false
            }

				row[i][value - 1] = true
				col[value - 1][j] = true
				boxes[xBoxes][yBoxes] = true
			
            
        }
    }
    return true
}

// Cần tối ưu