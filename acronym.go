package main

import (
	"fmt"
	"os"
	"bufio"
	"bytes"
	"strings"
)

func main() {
	fmt.Print("Enter phrase for its acronym: ")
	reader := bufio.NewReader(os.Stdin)
	phrase, _ := reader.ReadString('\n')
	words := strings.Split(phrase, " ")
	var buffer bytes.Buffer
	for _, word := range words {
		if word == "the" || word == "on" || word == "and" {
			// skip
		}else{
			buffer.Write([]byte(strings.ToUpper(string(word[0]))))
		}
	}
	fmt.Println(buffer.String())
}
