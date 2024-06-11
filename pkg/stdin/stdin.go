package stdin

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// StdIn ...
type StdIn struct {
	scanner *bufio.Scanner
}

// NewStdIn ...
func NewStdIn(split bufio.SplitFunc) *StdIn {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(split)
	return &StdIn{scanner}
}

// IsEmpty reports if the In is empty
func (s *StdIn) IsEmpty() bool {
	return !s.scanner.Scan()
}

// ReadString return next string by delimiter of ' '
func (s *StdIn) ReadString() string {
	return s.scanner.Text()
}

// ReadAllStrings returns all strings as a slice from the input
func ReadAllStrings() (words []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, strings.Split(line, " ")...)
	}

	return words
}

// ReadInt return next integer by delimiter of ' '
func (s *StdIn) ReadInt() int {
	s.scanner.Scan()
	v, _ := strconv.Atoi(s.scanner.Text())
	return v
}

// ReadAllInts returns all ints as an array
func (s *StdIn) ReadAllInts() (vals []int) {

	for s.scanner.Scan() {
		v, _ := strconv.Atoi(s.scanner.Text())
		vals = append(vals, v)
	}

	return vals
}
