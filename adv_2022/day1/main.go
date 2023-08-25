package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const fileName = "input_day_1.txt"

func main() {
  var mostCalories int
  var calories int
  caloriesList := []int{}

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
      caloriesList = append(caloriesList, calories)
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
  sort.Slice(caloriesList, func(i, j int) bool {
      return caloriesList[i] > caloriesList[j]
  })
  sum := 0
  for _, num := range caloriesList[:3] {
    fmt.Println(num)
    sum += num
  }
  fmt.Println("Answer: ", sum)
}
