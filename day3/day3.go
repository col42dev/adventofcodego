package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
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
	fmt.Printf("day3\n")
	lines, err := readLines("dirs.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	x:= 0;
	y:= 0;

	mom := map[int]map[int]int{ 0: map[int]int{0:42}}

	for _, line := range lines {
		for _, c := range line {
    		switch c {
    			 case 60: x-=1
    			 case 62: x+=1
    			 case 94: y+=1
    			 case 118: y-=1
    			 default: fmt.Println("default")
    		}
   
    		_, okx := mom[x]
			if !okx {
				mom[x] = make(map[int]int)
			}
			mom[x][y] = 42;


		}
	}

	totalCells := 0
	for i := range(mom) {
    	totalCells += len(mom[i])
	}

	fmt.Println(totalCells)
  
}

