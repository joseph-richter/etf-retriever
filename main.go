package main

import (
	"fmt"
)

func main() {
	link := "https://www.blackrock.com/fr/intermediaries/products/333675/ishares-short-duration-corp-bond-ucits-etf" // Remplace par l'URL que tu veux parser
	result := webContentGet(link)
	fmt.Println(result)
}
