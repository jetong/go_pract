package main

import "fmt"
import "strings"

func Operate(s []string, f func(st string) string) (result []string) {
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	list := []string{"dog", "cat", "rabbit"}
	result := Operate(list, Upper)
	fmt.Println(result)
}
