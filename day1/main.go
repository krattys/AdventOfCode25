package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const startPosition = 50

func main() {
	rotations := readInput()
	currentPosition := startPosition
	var res int
	for _, rotation := range rotations {
		direction := 1
		if rotation[0] == 'L' {
			direction = -1
		}
		magnitude, err := strconv.Atoi(rotation[1:])
		if err != nil {
			panic(err)
		}

		// Algebra
		res += (magnitude / 100)
		magnitude %= 100
		if magnitude == 0 {
			continue
		}
		nextPosition := (currentPosition + direction * magnitude)
		if (currentPosition > 0 && nextPosition <= 0) || nextPosition >= 100 {
			res++
		}
		currentPosition = (nextPosition + 100) % 100

		// Bruteforce
		// for range magnitude {
		// 	currentPosition += direction
		// 	if currentPosition == 0 || currentPosition == 100 {
		// 		res++
		// 	}
		// 	currentPosition = (currentPosition + 100) % 100
		// }
	}
	fmt.Println(res)
}

func readInput() []string {
	input, _ := io.ReadAll(os.Stdin)
	content := string(input)
	rotations := strings.Fields(content)
	return rotations
}
