package slice

import (
	"fmt"
	"sort"
	"strconv"
)

func Slice() []int {

	var s []int = make([]int, 3)
	var input string
	for {
		fmt.Println("Please enter an integer or X to exit! :")
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		s = append(s, value)
		sort.Ints(s[:])
		fmt.Println(s)

	}
	return s

}
