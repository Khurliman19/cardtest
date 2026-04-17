package main

import (
	"fmt"
	"testing"
)

func Example_maskCard() {
	card := "4111111111111111"
	fmt.Println(maskCard(card))

	// Output:
	// 4111 **** **** 1111
}

func Test_maskCard_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, but got none")
		}
	}()

	maskCard("123")
}
