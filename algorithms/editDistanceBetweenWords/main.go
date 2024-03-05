//write a function to returns whether two words are exactly one edit away
// one edit is:
// inserting one character anywhere in the word
// removing one character
// replacing one character

package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Printf("cats vs cat: %v\n", oneEditApart("cats", "cat"))
	fmt.Printf("cat vs at: %v\n", oneEditApart("cat", "at"))
	fmt.Printf("ca0 vs cat: %v\n", oneEditApart("cats", "cat"))
	fmt.Printf("cat vs dog: %v\n", oneEditApart("cat", "dog"))
	fmt.Printf("cat vs act: %v\n", oneEditApart("cat", "act"))
}

func oneEditApart(s1, s2 string) bool {

	//1. initial checks
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return false
	}
	//2. what if they are the same????
	var small, large string

	if len(s1) > len(s2) {
		large, small = s1, s2
	} else {
		large, small = s2, s1
	}

	//case1 check substring
	if len(s1) != len(s2) {
		return strings.Contains(large, small)
	}

	//case2 check every rune
	numberOfDiffs := 0
	for i, rune1 := range large {

		rune2 := rune(small[i])

		if rune1 != rune2 {
			numberOfDiffs++
		}

		if numberOfDiffs > 1 {
			return false
		}
	}

	return true
}
