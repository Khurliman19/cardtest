package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	cardInput := flag.String("card:", "", "номер карты")
	flag.Parse()

	card := strings.ReplaceAll(*cardInput, " ", "")
	card = strings.ReplaceAll(card, "-", "")

	if len(card) != 16 || !isDigits(card) {
		fmt.Println("Ошибка: неверный формат карты")
		return
	}

	fmt.Println("Нормализованный номер:", card)
	fmt.Println("Карта:", maskCard(card))
}

func maskCard(card string) string {
	if len(card) != 16 {
		panic("invalid card length")
	}
	return card[:4] + " **** **** " + card[12:]
}

func isDigits(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}
