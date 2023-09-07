package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const fileName = "./input_day_3.txt"

func mapItems(item string) map[rune]bool {
  var newMap map[rune]bool = make(map[rune]bool)
  for _, char := range item {
    newMap[char] = true
  }
  return newMap
}

func main() {
  file, _ := os.Open(fileName)
  defer file.Close()
  sc := bufio.NewScanner(file)

  var prioritiesSum int

  for sc.Scan() {
    firstItem := mapItems(sc.Text())
    sc.Scan()
    secondItem := mapItems(sc.Text())
    sc.Scan()
    thirdScan := mapItems(sc.Text())

    for firstElfItem := range firstItem {
      if secondItem[firstElfItem] && thirdScan[firstElfItem] {
        prioritiesSum += int(unicode.ToLower(firstElfItem) - 96)
        if unicode.IsUpper(firstElfItem) {
          prioritiesSum += 26
        }
        break
      }
    }
  }
  fmt.Println(prioritiesSum)
}
