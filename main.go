package main

import (
	"fmt"

	"joaopedroaat.me/leet-code/answers"
)

func main() {
	a := answers.RansomNote("a", "b")
	b := answers.RansomNote("aa", "ab")
	c := answers.RansomNote("aa", "aab")

	fmt.Printf("%t, %t, %t\n", a, b, c)
}
