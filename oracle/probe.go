package main

import (
	__json "encoding/json"
	__fmt "fmt"
)



import "math"

func order(n int) (result int) {
	for n != 0 {
		result++
		n /= 10
	}
	return result
}

// IsNumber returns true for an armstrong number
func IsNumber(n int) bool {
	originalNumber := n
	pow := order(n)
	sum := 0
	for n != 0 {
		remainder := n % 10
		sum += int(math.Pow(float64(remainder), float64(pow)))
		n /= 10
	}
	return originalNumber == sum
}

type R struct {
	Ok bool        `json:"ok"`
	V  interface{} `json:"v"`
}

func main() {
	inputs := []int{0, 5, 10, 153, 100, 9474, 9475, 9926315, 9926314}
	out := []R{}
	for _, x := range inputs {
		func() {
			defer func() { if r := recover(); r != nil { out = append(out, R{false, __fmt.Sprint(r)}) } }()
			out = append(out, R{true, IsNumber(x)})
		}()
	}
	b, _ := __json.Marshal(map[string]interface{}{"out": out})
	__fmt.Println(string(b))
}
