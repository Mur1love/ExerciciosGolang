package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Erro ao obter", url, err)
				return
			}
			defer resp.Body.Close()

			html, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Erro ao ler o corpo da resposta", err)
				return
			}

			r, err := regexp.Compile(`(?i)<title>(.*?)<\/title>`)
			if err != nil {
				fmt.Println("Erro ao compilar a expressão regular", err)
				return
			}

			matches := r.FindStringSubmatch(string(html))
			if len(matches) > 1 {
				c <- matches[1]
			} else {
				fmt.Println("Não foi possível encontrar o título para", url)
			}
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.google.com", "https://www.youtube.com")
	t2 := titulo("https://www.uol.com.br/", "https://www.twitch.com")

	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)
}
