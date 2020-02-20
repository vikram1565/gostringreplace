package stringreplace

import (
	"strings"
	"testing"
)

const (
	str = "Hello, My name is vikram. i like cricket and i am fan of dhoni."
)

var (
	s1 = strings.NewReplacer("vikram", "kunal")
	s2 = strings.NewReplacer("vikram", "kunal", "cricket", "football")
	s3 = strings.NewReplacer("vikram", "kunal", "cricket", "football", "dhoni", "messi")
	s4 = strings.NewReplacer("Hello", "Hi", "vikram", "kunal", "cricket", "football")
)

func BenchmarkReplaceOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replaceOne(str)
	}
}
func BenchmarkReplaceTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replaceTwo(str)
	}
}
func BenchmarkReplaceThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replaceThree(str)
	}
}
func BenchmarkReplaceFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replaceFour(str)
	}
}

func BenchmarkReplacerOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replacerOne(s1, str)
	}
}
func BenchmarkReplacerTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replacerTwo(s2, str)
	}
}
func BenchmarkReplacerThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replacerThree(s3, str)
	}
}
func BenchmarkReplacerFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replacerFour(s4, str)
	}
}
