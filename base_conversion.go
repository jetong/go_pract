// number base conversion (whole numbers only) for base 2 through 16

package main

import (
	"fmt"
	"math"
	"bytes"
	"strconv"
	"strings"
)

// convert num from base1 to base2
func changeBase(num string, base1 int, base2 int) string {
	return fromBase10(toBase10(num, base1), base2)
}

// convert to base 10
func toBase10(num string, base int) int {
	var result int 
	slice := []byte(num)
	for i,val := range slice {
		// convert each digit to its actual numerical value representation
		if val >= 65 {	// A-F
			val -= 55
		}else{			// 0-9
			val -= 48
		}
		// ex: 1A5 (base16) becomes 1*16^2 + 10*16^1 + 5*16^0 = 421 (base10)
		result += int(val)*int(math.Pow(float64(base), float64(len(slice)-1-i)))
	}
	return result
}

// convert from base 10 to target base
func fromBase10(num, base int) string {
	if base == 10 {
		return strconv.Itoa(num)	// no conversion needed
	}
	var i, remainder int
	var buffer bytes.Buffer
	for base <= num {
		remainder = int(math.Mod(float64(num), float64(base)))
		// push remainders into the buffer
		buffer.Write([]byte(toLetter(remainder)))
		num = num / base
		i++
	}
	// once base is too large to divide num, we've finished the loop
	// but we still need the last value of num added to the buffer
	buffer.Write([]byte(toLetter(num)))
	s := buffer.String()
	reverseString(&s)
	return s
}

// use double value assignemnt on slice elements to reverse string
func reverseString(str *string) {
	s := strings.Split(*str, "")
	for i:=0; i<=len(s)/2; i++ {
		s[i],s[len(s)-1-i] = s[len(s)-1-i],s[i]
	}
	*str = strings.Join(s, "")
}

// use ASCII to determine letters A-F or 0-9
func toLetter(num int) string {
	if num > 9 {
		return string(num + 55)		// A-F  'A'=65
	}else{
		return string(num + 48)		// 0-9  '0'=48
	}
}


func main() {
	var num string
	var base1, base2 int
	fmt.Print("Enter number to convert: ")
	fmt.Scanln(&num)
	fmt.Print("Enter start base: ")
	fmt.Scanln(&base1)
	fmt.Print("Enter end base: ")
	fmt.Scanln(&base2)

	fmt.Printf("%v base %v is %v in base %v\n", num, base1, changeBase(num, base1, base2), base2)
}
