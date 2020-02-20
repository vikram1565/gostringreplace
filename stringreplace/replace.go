package stringreplace

import (
	"fmt"
	"strings"
)

// InitReplace - init all functions
func InitReplace() {
	// constant string
	str := "Hello, My name is vikram. i like cricket and i am fan of dhoni."
	// function calling
	fmt.Println("replaceOne: ", replaceOne(str))
	fmt.Println("replaceTwo: ", replaceTwo(str))
	fmt.Println("replaceThree: ", replaceThree(str))
	fmt.Println("replaceFour: ", replaceFour(str))
}

// replaceOne - replace one substring from given string
func replaceOne(str string) string {
	s := strings.Replace(str, "vikram", "kunal", 1)
	return s
}

// replaceTwo - replace two substring from given string
func replaceTwo(str string) string {
	s := strings.Replace(str, "vikram", "kunal", 1)
	s = strings.Replace(s, "cricket", "football", 1)
	return s
}

// replaceThree - replace three substring from given string
func replaceThree(str string) string {
	s := strings.Replace(str, "vikram", "kunal", 1)
	s = strings.Replace(s, "cricket", "football", 1)
	s = strings.Replace(s, "dhoni", "messi", 1)
	return s
}

// replaceFour - replace four substring from given string
func replaceFour(str string) string {
	s := strings.Replace(str, "Hello", "Hi", 1)
	s = strings.Replace(s, "vikram", "kunal", 1)
	s = strings.Replace(s, "cricket", "football", 1)
	s = strings.Replace(s, "dhoni", "messi", 1)
	return s
}
