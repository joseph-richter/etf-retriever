module etf-retriever/main

go 1.24.3

replace etf-retriever/website_treatment => ./website_treatment.go

require github.com/PuerkitoBio/goquery v1.10.3

require (
	github.com/andybalholm/cascadia v1.3.3 // indirect
	golang.org/x/net v0.39.0 // indirect
)
