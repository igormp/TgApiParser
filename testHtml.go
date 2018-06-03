package main

import (
	"fmt"
	"golang.org/x/net/html"
	_ "io/ioutil"
	"log"
	"net/http"
	_ "os"
	_ "strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func findTokenByData(s string) html.Token {

}

func main() {
	resp, err := http.Get("https://core.telegram.org/bots/api")
	check(err)
	defer resp.Body.Close()
	//b, err := ioutil.ReadAll(resp.Body)
	//check(err)

	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()
		done := false
		switch tt {
		case html.ErrorToken:
			// End of the document, we're done
			return

		case html.StartTagToken:
			t := z.Token()

			if t.Data == "h3" {
				for {
					z.Next()
					a := z.Token()
					if a.Type.String() == "Text" && a.Data == "Available types" {
						fmt.Println(a.Data)
						done = true
					}
					if done {
						break
					}
				}
				if done {
					break
				}
				/*z.Next()
				z.Next()
				z.Next()
				z.Next()
				z.Next()
				a := z.Token()
				fmt.Println(a.Type)
				fmt.Println(a.Data)*/
				/*for {
					st := z.Next()
					if st == html.TextToken {
						a := z.Token()
						fmt.Println(a.Data)
						break
					}
				}*/
			}
			/*if isTopic {
				z.Next()
				z.Next()
				z.Next()
				z.Next()
				fmt.Println(z.Next())
				fmt.Println("    ", string(z.Raw()))
				fmt.Println("    Found H4")
				fmt.Println(t.Type)
				fmt.Println(len(t.Attr))
				for _, a := range t.Attr {
					if a.Key == "href" {
						fmt.Println("Found href:", a.Val)
					}
				}
			}*/
		}
	}

}
