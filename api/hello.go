package handler

import (
	"fmt"
	"github.com/go-rod/rod"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")

	url := "https://en.wikipedia.org/wiki/The_Lord_of_the_Rings"
	browser := rod.New().MustConnect()
	defer browser.Close()

	page := browser.MustPage(url).MustWaitLoad()
	bodyElements, err := page.Elements(".infobox.vcard")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	for _, element := range bodyElements {
		text, _ := element.Text()
		fmt.Println(text)
	}
}
