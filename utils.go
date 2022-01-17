package advent2021

import (
	"strconv"
	"strings"
)

func StringToIntSlice(input string) []int {
	var intSlice []int
	for _, str := range strings.Split(strings.TrimSpace(input), "\n") {
		converted, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, converted)
	}
	return intSlice
}
