package main

import (
	"fmt"
	"math"
	"bytes"
	"strconv"
	"strings"
)

func reverseString(str *string) {
	s := strings.Split(*str, "")
	for i:=0; i<=len(s)/2; i++ {
		s[i],s[len(s)-1-i] = s[len(s)-1-i],s[i]
	}
	*str = strings.Join(s, "")
}

func toBase10(num string, base int) int {
	var result int 
	slice := []byte(num)
	for i,val := range slice {
		if val >= 65 {	// A-F
			val -= 55
		}else{			// 0-9
			val -= 48
		}
		result += int(val)*int(math.Pow(float64(base), float64(len(slice)-1-i)))
	}
	return result
}

func fromBase10(num, base int) string {
	if base == 10 {
		return strconv.Itoa(num)
	}
	var i, remainder int
	var buffer bytes.Buffer
	for base <= num {
		remainder = int(math.Mod(float64(num), float64(base)))
		buffer.Write([]byte(toLetter(remainder)))
		num = num / base
		i++
	}
	buffer.Write([]byte(toLetter(num)))
	s := buffer.String()
	reverseString(&s)
	return s
}

func toLetter(num int) string {
	if num > 9 {
		return string(num + 55)		// 10 + 55 = 65 -> 'A'
	}else{
		return string(num + 48)		// 0 + 48 = 48 -> '0'
	}
}

// change num from base1 to base2 representation
func changeBase(num string, base1 int, base2 int) string {
	return fromBase10(toBase10(num, base1), base2)
}

func main() {
	var num string
	var base1, base2 int
	fmt.Print("Enter number: ")
	fmt.Scanln(&num)
	fmt.Print("Enter start base: ")
	fmt.Scanln(&base1)
	fmt.Print("Enter end base: ")
	fmt.Scanln(&base2)

	fmt.Println(changeBase(num, base1, base2))
}
