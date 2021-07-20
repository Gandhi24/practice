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

func main() {

}
