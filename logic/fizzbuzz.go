package logic

import (
	"strconv"
)

func Fizzbuzz() (result []string) {
	var input = 60
	for i := 1; i < input; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		}
		if i%3 != 0 && i%5 == 0 {
			result = append(result, "Buzz")
		}
		if i%3 == 0 && i%5 != 0 {
			result = append(result, "Fizz")
		}
		if i%3 != 0 && i%5 != 0 {
			result = append(result, strconv.Itoa(i))
		}
	}
	return
}
