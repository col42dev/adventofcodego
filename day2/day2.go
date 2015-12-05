package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "regexp"
  "strconv"
  "sort"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	fmt.Printf("day2 part 1\n")
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var re, errRegExp = regexp.Compile("(.*)x(.*)x(.*)")

	if errRegExp != nil {
		log.Fatalf("regexp err: %s", errRegExp)
	}

	var totalArea int = 0

	for _, line := range lines {

		result := re.FindStringSubmatch(line)
		if (result != nil) {
			var l, w, h int
			l, err = strconv.Atoi(result[1])
			w, err = strconv.Atoi(result[2])
			h, err = strconv.Atoi(result[3])
			totalArea += 2*l*w + 2*w*h + 2*h*l;

			sides := []int{l, w, h}
			sort.Ints(sides)
			totalArea += sides[0] * sides[1]
		}

  	}

  	fmt.Printf( " puzzle answer = %d\n", totalArea)

  	fmt.Printf("day2 part 2\n")

	var totalRibbonLength int = 0
	
	for _, line := range lines {

		result := re.FindStringSubmatch(line)
		if (result != nil) {
			var l, w, h int
			l, err = strconv.Atoi(result[1])
			w, err = strconv.Atoi(result[2])
			h, err = strconv.Atoi(result[3])


			sides := []int{l, w, h}
			sort.Ints(sides)
			totalRibbonLength += sides[0] * 2  +  sides[1] * 2
			totalRibbonLength += l * w * h;
		}

  	}

  	fmt.Printf( " puzzle answer = %d\n", totalRibbonLength)
  
}

