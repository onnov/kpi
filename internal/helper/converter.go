package helper

import "strconv"

// Sti string to int
func Sti(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = 0
	}

	return i
}
