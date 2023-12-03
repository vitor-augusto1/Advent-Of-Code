package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const FILE_NAME = "adventofcode.com_2023_day_1_input.txt"

func findNumbers(s string) []int64 {
  var sliceOfSums []int64
  re := regexp.MustCompile("[0-9]+")
  matches := re.FindAllString(s, -1)
  if len(matches) != 0 {
    firstValue := matches[0]
    firstCharacterInFirstValue := firstValue[0]
    lastNumber := matches[len(matches)-1]
    lastCharacterInLastNumber := lastNumber[len(lastNumber)-1]
    string1 := string(firstCharacterInFirstValue)
    string2 := string(lastCharacterInLastNumber)
    concatenated := fmt.Sprintf("%s%s", string1, string2)
    convertedStringToInteger64, _ := strconv.ParseInt(concatenated, 0, 64)
    sliceOfSums = append(sliceOfSums, convertedStringToInteger64)
    return sliceOfSums
  }
  return nil
}

func main() {
  fileContent, _ := os.ReadFile(FILE_NAME)
  fileContentToString := string(fileContent)
  splitedString := strings.Split(fileContentToString, "\n")
  var tmpSum int64
  var sliceOfNumbers []int64
  for _, line := range splitedString {
    slc := findNumbers(line)
    if slc == nil {
      continue
    }
    sliceOfNumbers = append(sliceOfNumbers, slc...)
  }
  for _, number := range sliceOfNumbers {
    tmpSum += number
  }
  fmt.Println("RESULT :::: ", tmpSum)
}
