package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func scrape() {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", "https://www.google.com/search?q=gbp+to+php", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Chrome")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP body. ", err)
		os.Exit(1)
	}

	// Debug
	// fmt.Println(string(body))

	re := regexp.MustCompile("1 Pound sterling =\\s\\d+\\.\\d+")
	found := re.FindString(string(body))
	if len(found) < 0 {
		fmt.Println("Error: no exchange rate found")
		os.Exit(1)
	}
	fmt.Println("GBP to PHP", strings.Split(found, "1 Pound sterling = ")[1])
	fmt.Println("---")
	fmt.Println("Transfer Money (World Remit) | href=https://www.worldremit.com/en/Philippines")
}

func main() {
	scrape()
}
