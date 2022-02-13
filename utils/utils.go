package utils

import (
	"strconv"
)

func StringSliceToIntSlice(input []string) []int {
	var intSlice []int
	for _, str := range input {
		converted, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, converted)
	}
	return intSlice
}
