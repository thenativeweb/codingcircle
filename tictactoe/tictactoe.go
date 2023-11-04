package tictactoe

type Field struct {
	Row    int
	Column int
}

func HasWon(fields []Field, size int) bool {
	rowSums := make([]int, size)
	columnSums := make([]int, size)

	slashDiagonalSum := 0
	backslashDiagonalSum := 0

	for _, field := range fields {
		rowSums[field.Row]++
		columnSums[field.Column]++

		if field.Row == field.Column {
			slashDiagonalSum++
		}
		if field.Row == size-field.Column-1 {
			backslashDiagonalSum++
		}

		if rowSums[field.Row] == size ||
			columnSums[field.Column] == size ||
			slashDiagonalSum == size ||
			backslashDiagonalSum == size {
			return true
		}
	}

	return false
}
