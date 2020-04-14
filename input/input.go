package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//SliceInt take input argumants into slises type int
func SliceInt(sizeSlice int) []int {
	var values []int
	var inputNumber []string
	count := 0
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		inputNumber = append(inputNumber, input.Text())
		value, err := strconv.Atoi(inputNumber[count])
		if err != nil {
			fmt.Println(err, "please input a valid number")
			continue
		} else {
			values = append(values, value)
			count++
		}
		if count >= sizeSlice {
			break
		}
	}
	return values
}
