package main
import "fmt"

func big_sum(int_array []int64) int64 {
	var total_sum int64
	for _, big_int := range int_array {
		total_sum += big_int
	}
	return total_sum
}

func get_big_int_array(array_sz int, int_array []int64) []int64 {
	var current_int int64
	for i := 0; i < array_sz; i++ {
		fmt.Scanf("%v", &current_int)
		int_array = append(int_array, current_int)
	}
	return int_array
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var array_sz int
	var int_array []int64
	fmt.Scanf("%v", &array_sz)
	int_array = get_big_int_array(array_sz, int_array)
	sum := big_sum(int_array)
	fmt.Printf("%v\n", sum)
}