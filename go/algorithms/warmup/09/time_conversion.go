package main
import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func convertTime(twelve_hour string) string {
	splitTime := strings.Split(twelve_hour, ":")
	hours := splitTime[0]
	minutes := splitTime[1]
	seconds := splitTime[2][:2]
	time_meridiem := splitTime[2][2:4] // to chop off '\n'
	
	int_hours, _ := strconv.Atoi(hours)
	int_hours = int_hours % 12
	if time_meridiem == "PM" {
		int_hours = int_hours + 12
	}
	hours = strconv.Itoa(int_hours)
	// pad hour back to 2 digits
	if len(hours) < 2 {
		hours = fmt.Sprintf("0%s", hours)
	}
	twenty_four_hour := fmt.Sprintf("%s:%s:%s", hours, minutes, seconds)
	return twenty_four_hour
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReader(os.Stdin)
	twelve_hour, _ := reader.ReadString('\n')
	twenty_four_hour := convertTime(twelve_hour)
	fmt.Printf(twenty_four_hour)
}