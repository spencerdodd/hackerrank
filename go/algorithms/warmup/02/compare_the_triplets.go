package main
import "fmt"

func compareTriplets(t1 [3]int, t2 [3]int) (int, int) {
	t1_score := 0
	t2_score := 0
	for i := 0; i < 3; i++ {
		if t1[i] > t2[i] {
			t1_score += 1
		} else if t1[i] < t2[i] {
			t2_score += 1
		}
	}
	return t1_score, t2_score
}

func parseTriplet(triplet [3]int) [3]int {
	var val int
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &val)
		triplet[i] = val
	}
	return triplet
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t1, t2 [3]int
	t1 = parseTriplet(t1)
	t2 = parseTriplet(t2)
	t1_score, t2_score := compareTriplets(t1, t2)
	fmt.Printf("%d %d\n", t1_score, t2_score)
}