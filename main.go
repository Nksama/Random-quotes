package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Qt struct {
	Content string
	Author  string
}

func quote() {

	for true {
		resp, err := http.Get("https://api.quotable.io/random")

		if err != nil {
			println(err)
		}

		quotes, _ := ioutil.ReadAll(resp.Body)
		var quote Qt
		json.Unmarshal([]byte(quotes), &quote)
		fmt.Printf("%v \n - %v \n \n", quote.Content, quote.Author)
		time.Sleep(5 * time.Second)

	}

}

func main() {
	quote()
}
