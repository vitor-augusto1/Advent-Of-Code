
package main

// Rock: A, X
// Paper: B, Y
// Scissors: C, Z
// Total Score is the sum of your scores for each round
// Single round Score: Shape (1 rock, 2 paper, 3 scissors) + outcome (0 lost, 3 draw, 6 win)

import (
	"bufio"
	"fmt"
	"os"
)

const fileNames = "./input_day_2.txt"

func main() {
	var score int
  file, _ := os.Open(fileNames)
  defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		round := sc.Text()

		switch round[2:] {
		case "Y":
			score += 3
		case "Z":
			score += 6
		}

		switch round {
		case "C Z", "A Y", "B X":
			score += 1
		case "A Z", "B Y", "C X":
			score += 2
		case "B Z", "C Y", "A X":
			score += 3
		}
	}

	fmt.Println(score)
}
