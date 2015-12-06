package main

import (
	"fmt"
	"bufio"
  	"log"
  	"os"
	"strings"
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

func day5pt1() {

	fmt.Printf("--- Day 5: Doesn't He Have Intern-Elves For This? ---\n")
	fmt.Printf("\n")
	fmt.Printf("Santa needs help figuring out which strings in his text file are naughty or nice.\n")
	fmt.Printf("\n")
	fmt.Printf("A nice string is one with all of the following properties:\n")
	fmt.Printf("\n")
	fmt.Printf("It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.\n")
	fmt.Printf("It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).\n")
	fmt.Printf("It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.\n")
	fmt.Printf("For example:\n")
	fmt.Printf("\n")
	fmt.Printf("ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.\n")
	fmt.Printf("aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.\n")
	fmt.Printf("jchzalrnumimnmhp is naughty because it has no double letter.\n")
	fmt.Printf("haegwjzuvuyypxyu is naughty because it contains the string xy.\n")
	fmt.Printf("dvszwmarrgswjxmb is naughty because it contains only one vowel.\n")
	fmt.Printf("How many strings are nice?\n")
	fmt.Printf("")


	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

 	vowels := []string{"a", "e", "i", "o", "u"} 

	badwords := [4]string{"ab", "cd", "pq", "xy"}

	var niceStringCount int = 0

	for _, line := range lines {

		
		// badwords
		var hasBadWord bool = false
		for _, badword := range badwords {
			if strings.Contains(line, badword) {
				hasBadWord = true
			}
		}
		
		// vowels
		var hasToFewVowels bool = false
		var vowelCount int = 0
		for _, vowel := range vowels {
			vowelCount += strings.Count(line, vowel)
		}
		if vowelCount < 3 {
			hasToFewVowels = true
		}
		//fmt.Printf(" %s : vowelCount %d %t\n", line, vowelCount, hasToFewVowels)


		// repeats
		var lastLetter rune = 0
		var norepeats bool = true
		for _, letter := range line {
			if lastLetter == letter {
				norepeats = false;
			}
			lastLetter = letter;
		}
		//fmt.Printf(" %s : norepeats %t\n", line, norepeats)

		if hasBadWord || hasToFewVowels || norepeats {
			
		} else {
			niceStringCount += 1
		}


		//fmt.Printf(" %s : hasBadWord %t, hasToFewVowels %t, norepeats %t\n", line, hasBadWord, hasToFewVowels, norepeats)
		//255
	}
	fmt.Println(len(lines), niceStringCount)
 
}
/*
func main() {
	fmt.Printf("--- Part Two ---\n")
	fmt.Printf("\n")
	fmt.Printf("Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice. None of the old rules apply, as they are all clearly ridiculous.\n")
	fmt.Printf("\n")
	fmt.Printf("Now, a nice string is one with all of the following properties:\n")
	fmt.Printf("\n")
	fmt.Printf("It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).\n")
	fmt.Printf("It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.\n")
	fmt.Printf("For example:\n")
	fmt.Printf("\n")
	fmt.Printf("qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).\n")
	fmt.Printf("xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.\n")
	fmt.Printf("uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.\n")
	fmt.Printf("ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.\n")
	fmt.Printf("How many strings are nice under these new rules?\n")


	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

 


}	*/

