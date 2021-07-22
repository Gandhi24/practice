package main

type Matrix struct {
	rows, cols int
	grid       [][]int
}

func (m Matrix) Rows() int {
	return m.rows
}

func (m Matrix) Cols() int {
	return m.cols
}

func (m *Matrix) setElement(i, j, val int) {
	m.grid[i][j] = val
}

func (m *Matrix) add(other *Matrix) {
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			m.grid[row][col] += other.grid[row][col]
		}
	}
}

func main() {
	var m *Matrix = &Matrix{
		rows: 2,
		cols: 4,
		grid: [][]int{
			{0, 1, 2, 3},
			{4, 5, 6, 7},
		},
	}

	m.setElement(0, 0, 100)
}
