package stringreplace

import (
	"fmt"
	"strings"
)

// InitReplacer - init all functions
func InitReplacer() {
	// constant string
	str := "Hello, My name is vikram. i like cricket and i am fan of dhoni."
	// New Replacers
	s1 := strings.NewReplacer("vikram", "kunal")
	s2 := strings.NewReplacer("vikram", "kunal", "cricket", "football")
	s3 := strings.NewReplacer("vikram", "kunal", "cricket", "football", "dhoni", "messi")
	s4 := strings.NewReplacer("Hello", "Hi", "vikram", "kunal", "cricket", "football")
	// function calling
	fmt.Println("replacerOne: ", replacerOne(s1, str))
	fmt.Println("replacerTwo: ", replacerTwo(s2, str))
	fmt.Println("replacerThree: ", replacerThree(s3, str))
	fmt.Println("replacerFour: ", replacerFour(s4, str))
}

// replacerOne - replace one substring from given string
func replacerOne(s *strings.Replacer, str string) string {
	return s.Replace(str)
}

// replacerTwo - replace two substring from given string
func replacerTwo(s *strings.Replacer, str string) string {
	return s.Replace(str)
}

// replacerThree - replace three substring from given string
func replacerThree(s *strings.Replacer, str string) string {
	return s.Replace(str)
}

// replacerFour - replace four substring from given string
func replacerFour(s *strings.Replacer, str string) string {
	return s.Replace(str)
}
