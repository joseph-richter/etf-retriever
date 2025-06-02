package main

import (
	"fmt"
)

func main() {
	//url[0] = "https://www.blackrock.com/fr/intermediaries/products/333675/ishares-short-duration-corp-bond-ucits-etf" // Remplace par l'URL que tu veux parser
	//url[1] = "https://www.blackrock.com/fr/intermediaries/products/251843/ishares-euro-high-yield-corporate-bond-ucits-etf"
	url := "https://www.blackrock.com/fr/intermediaries/products/333675/ishares-short-duration-corp-bond-ucits-etf"
	result := webContentGet(url)
	fmt.Println(result)
}
