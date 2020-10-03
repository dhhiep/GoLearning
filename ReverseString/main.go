package main

import "fmt"

func main() {
	fmt.Printf("Hello, World!!! will be reversed to %v\n", reverseString("Hello, World!!!"))
}

func reverseString(s string) string {
	r := []rune(s)

	for i := 0; i <= len(r)/2; i++ {
		if i >= len(r)/2 {
			break
		}

		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}

	return string(r)
}
