package numbers

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed numbers.txt
var data []byte

func Sum() int {
	lines := strings.Split(string(data), "\r\n")
	var sum int
	for _, line := range lines {
		if line != "" {
			val, _ := strconv.Atoi(line)
			sum += val
		}
	}
	return sum
}
