package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := [][]string{}

	// 30 valid pair of ramdom strings to concatenate
	for i := 0; i < 30; i++ {
		value := []string{z01.RandASCII(), z01.RandASCII()}
		table = append(table, value)
	}
	table = append(table,
		[]string{"Hello!", " How are you?"},
	)
	for _, arg := range table {
		z01.Challenge("Concat", student.Concat, correct.Concat, arg[0], arg[1])
	}
}
