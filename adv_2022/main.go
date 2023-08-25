package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "input_day_1.txt"

func main() {
  var mostCalories int
  var calories int
  fileContent, err := os.ReadFile(fileName)
  if err != nil {
    log.Fatal(err)
    return
  }
  fileContentToString := string(fileContent)
  splitedString := strings.Split(fileContentToString, "\n")
  for _, line := range splitedString {
    if len(line) == 0 {
      if calories > mostCalories {
        mostCalories = calories
      }
      calories = 0
      continue
    }
    intCalorie, err := strconv.Atoi(string(line))
    if err != nil {
      log.Fatal("Atoi error: ", err)
      return
    }
    calories += intCalorie
  }
  fmt.Println(mostCalories)
}
