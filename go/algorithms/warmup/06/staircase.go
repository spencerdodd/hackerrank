package main
import (
	"bytes"
	"fmt"
)

func makeStaircase(staircase_sz int) {
	var stair_buf bytes.Buffer
	for i := 0; i < staircase_sz; i++ {
		last_space := staircase_sz - i - 1
		for j := 0; j < staircase_sz; j++ {
			if j < last_space {
				stair_buf.WriteString(" ")
			} else {
				stair_buf.WriteString("#")
			}
		}
		stair_buf.WriteString("\n")
	}
	fmt.Println(stair_buf.String())
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var staircase_sz int
	fmt.Scanf("%v", &staircase_sz)
	makeStaircase(staircase_sz)
}