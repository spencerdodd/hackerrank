package main
import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func minMaxPermutations(input_ints []int) (int, int) {
	var sum, min, max int = 0, 0xffffffff, 0
	for i := 0; i < len(input_ints); i++ {
		sum = 0
		for j := 0; j < len(input_ints); j++ {
			if j != i {
				sum += input_ints[j]
			}
		}
		if sum < min {
			min = sum
		}
		if sum > max {
			max = sum
		}
	}
	return min, max
}

func stringToSlice(input_str string) []int {
	str_slice := strings.Fields(input_str)
	var int_slice []int
	for i := 0; i < len(str_slice); i++ {
		current_int, _ := strconv.Atoi(str_slice[i])
		int_slice = append(int_slice, current_int)
	}
	return int_slice
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// storing the ints from input
	reader := bufio.NewReader(os.Stdin)
	input_str, _ := reader.ReadString('\n')
	input_ints := stringToSlice(input_str)
	min, max := minMaxPermutations(input_ints)
	fmt.Printf("%d %d", min, max)
}