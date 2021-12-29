package quotegenerator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const quoteApiUrl = "https://quotes.rest/qod?language=en"

type QuoteGenerator struct{}

func (q *QuoteGenerator) GetNewQuote() (string, error) {
	body, err := q.httpGet(quoteApiUrl)
	if err != nil {
		return "", err
	}

	quote, err := q.extractQuoteFromHTTPBody(body)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return quote, nil
}

func (q *QuoteGenerator) httpGet(url string) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to HTTP GET (status: %s): %s", r.Status, err)
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read HTTP GET body: %s", err)
	}

	if len(string(body)) == 0 {
		return "", fmt.Errorf("empty HTTP GET body: %s", err)
	}

	return string(body), nil
}

func (q *QuoteGenerator) extractQuoteFromHTTPBody(body string) (quote string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(
				"failed to extract quote. Response body structure has either change or the request limit has been reached")
		}
	}()

	// Error prone way to unpack body but it doesn't require creating a complex struct
	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return
	}

	// Drill down the body and extract the quote
	contents := data["contents"].(map[string]interface{})
	quotes := contents["quotes"].([]interface{})
	quoteContent := quotes[0].(map[string]interface{})

	quote = fmt.Sprintf("%s - %s", quoteContent["quote"].(string), quoteContent["author"].(string))

	return
}
