package main

// Rock: A, X
// Paper: B, Y
// Scissors: C, Z
// Total Score is the sum of your scores for each round
// Single round Score: Shape (1 rock, 2 paper, 3 scissors) + outcome (0 lost, 3 draw, 6 win)

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const fileName = "./input_day_2.txt"

//const stringExample = `
//A Y
//B X
//C Z
//`

func main() {
  var totalScore int
  scores := map[string]int{
    "A X": 4,
    "A Y": 8,
    "A Z": 3,
    "B X": 1,
    "B Y": 5,
    "B Z": 9,
    "C X": 7,
    "C Y": 2,
    "C Z": 6,
  }
  file, err := os.Open(fileName)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  fileScanner := bufio.NewScanner(file)
  for fileScanner.Scan() {
    roundPlay := fileScanner.Text()
    totalScore += scores[roundPlay]
  }
  fmt.Println("Total score: ", totalScore)
}
