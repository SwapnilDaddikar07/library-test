package string_utils

import "fmt"

func PrintLength(s string) {
	fmt.Println("Length is ", len(s))
}

func IsValid(s string) bool {
	return len(s) > 0
}
