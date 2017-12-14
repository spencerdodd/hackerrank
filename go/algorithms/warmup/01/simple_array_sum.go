package main
import "fmt"

func sumArray(array_sz uint32) uint32 {
	var i, sum, current_val uint32
	sum = 0
	for i = 0; i<array_sz; i++ {
		fmt.Scanf("%v", &current_val)
		sum += current_val
	}
	return (sum)
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var array_sz, res uint32
	fmt.Scanf("%v", &array_sz)
	res = sumArray(array_sz)
	fmt.Println(res)
}