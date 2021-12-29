package main

import (
	quotegenerator "CurtisM132/FortuneClone/QuoteGenerator"
	"fmt"
)

func main() {
	q := quotegenerator.QuoteGenerator{}

	quote, err := q.GetNewQuote()
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Quote of the Day: ", quote)
}
