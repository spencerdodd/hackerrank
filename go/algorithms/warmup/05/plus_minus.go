package main
import "fmt"

func fractionateArray(array_sz int) (float64, float64, float64) {
	var pos, neg, zero, total float64
	var current int
	for i := 0; i < array_sz; i++ {
		fmt.Scanf("%v", &current)
		total += 1
		if current > 0 {
			pos += 1
		}
		if current < 0 {
			neg += 1
		}
		if current == 0 {
			zero += 1
		}
	}
	var pos_frac, neg_frac, zero_frac float64
	pos_frac = pos / total
	neg_frac = neg / total
	zero_frac = zero / total

	return pos_frac, neg_frac, zero_frac
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var array_sz int
	fmt.Scanf("%v", &array_sz)
	pos, neg, zero := fractionateArray(array_sz)
	fmt.Printf("%v\n%v\n%v\n", pos, neg, zero)
}