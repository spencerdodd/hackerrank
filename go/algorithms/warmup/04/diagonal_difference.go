package main
import "fmt"
import "math"

func diagonalDifference(mat_sz int) int {
	var current int
	// slice with len 0 (no current elements) and capacity mat_sz
	matrix := make([][]int, 0, mat_sz)

	// fill/create the matrix
	for i := 0; i < mat_sz; i++ {
		// slice with len 0 and capacity mat_sz
		row := make([]int, 0, mat_sz)
		for j := 0; j < mat_sz; j++ {
			fmt.Scanf("%v", &current)
			row = append(row, current)
		}
		matrix = append(matrix, row)
	}

	// compute the top-left -> bottom-right diagonal
	var tl_sum int
	for i := 0; i < mat_sz; i++ {
		tl_sum += matrix[i][i]
	}

	// compute the top-right -> bottom-left diagonal
	var tr_sum int
	for i := 0; i < mat_sz; i++ {
		tr_sum += matrix[i][mat_sz - 1 - i]
	}

	total_dif := float64(tl_sum - tr_sum)
	// get abs value
	abs := int(math.Abs(total_dif))

	return (abs)
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var mat_sz int
	fmt.Scanf("%v", &mat_sz)
	res := diagonalDifference(mat_sz)
	fmt.Printf("%d", res)
}