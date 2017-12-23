package main
import "fmt"

type tallest struct {
	height int
	count int
}

func updatePeak(peak *tallest, new_candle int) {
	if new_candle > peak.height {
		peak.height = new_candle
		peak.count = 1
	} else if new_candle == peak.height {
		peak.count++
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	peak := tallest{0,0}
	var num_candles, new_candle int
	fmt.Scanf("%v", &num_candles)
	for i := 0; i < num_candles; i++ {
		fmt.Scanf("%v", &new_candle)
		updatePeak(&peak, new_candle)
	}
	fmt.Printf("%d\n", peak.count)
}