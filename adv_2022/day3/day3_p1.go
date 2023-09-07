package main

import (
	"bufio"
	"fmt"
	"unicode"
	"os"
)

const fileName = "./input_day_3.txt"

func main() {
  file, _ := os.Open(fileName)
  defer file.Close()
  sc := bufio.NewScanner(file)

  var prioritiesSum int

  for sc.Scan() {
    var newMap map[rune]rune = make(map[rune]rune)
    currentLine := sc.Text()
    middleString := len(currentLine) / 2
    leftSubstring := currentLine[:middleString]
    rightSubstring := currentLine[middleString:]
    for _, charLeft := range leftSubstring {
      currentCharacter := charLeft
      newMap[currentCharacter] = currentCharacter
    }
    for _, charRight := range rightSubstring {
      character, ok := newMap[charRight]
      if ok {
        prioritiesSum += int(unicode.ToLower(character) - 96)
        if unicode.IsUpper(character) {
          prioritiesSum += 26
        }
        break
      }
    }
  }
  fmt.Println(prioritiesSum)
}
