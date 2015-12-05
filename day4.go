package main

import (
	"fmt"
	"crypto/md5"
	"strconv"
	"encoding/hex"
	"strings"
)

var prefix = "yzbqklnj"

func day4(l int) int {
	res := 0
	target := strings.Repeat("0", l)
	for i := 0; i < 9999999; i++ {	
			m := md5.Sum([]byte(prefix + strconv.Itoa(i)))
			hex := hex.EncodeToString(m[:])
			if strings.HasPrefix(hex, target) {
				res = i
				break
			}
	}

	return res
}

func main() {

   	fmt.Printf("day4\n")
	fmt.Printf("%d zeros: %d\n", 5, day4(5))
	fmt.Printf("%d zeros: %d\n", 6, day4(6))
  
}

