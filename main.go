package main

import (
	"fmt"
	"strings"

	hitopia "github.com/sonnyariady/go-modul-hitopia"
)

func main() {

	objbracket := hitopia.BalancedBracket{Input: "[123)]"}
	fmt.Println(objbracket.AreBracketsBalanced())

	objsw := hitopia.StringWeight{InputString: "abbcccd", ArrayQuery: "1,3,9,8"}
	objsw.GenerateResult()

	if objsw.IsValid() {
		fmt.Println(strings.Join(objsw.ArrQueryResult, ","))
	} else {
		fmt.Println(strings.Join(objsw.ListInvalid, ","))
	}

	objLargestPalindrom := hitopia.LargestPalindrom{InputString: "3943", K: 1}
	fmt.Println("Largest Palindrom: ", objLargestPalindrom.GenerateResult())

	objLargestPalindrom.InputString = "932239"
	objLargestPalindrom.K = 2

	fmt.Println("Largest Palindrom: ", objLargestPalindrom.GenerateResult())
}
