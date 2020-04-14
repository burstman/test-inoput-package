package main

import (
	"fmt"
	"sort"

	"minmax/input"
)

func main() {
	values := input.SliceInt(6)
	min(values...)
	fmt.Println(values)
	max(values...)
	fmt.Println(values)
}
func min(Num ...int) {
	sort.Ints(Num)
}
func max(Num ...int) {
	sort.Sort(sort.Reverse(sort.IntSlice(Num)))
}
