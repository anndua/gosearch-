package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gosearch <query>")
		return

	}

	queryparts := os.Args[1:]
	query := strings.Join(queryparts, "")

	params := url.Values{}
	params.Set("q", query)
	searchURL := "https://www.google.com/search?" + params.Encode()
	cmd := exec.Command("/mnt/c/Program Files/Google/Chrome/Application/chrome.exe", searchURL)
	err := cmd.Run()
	if err != nil {
		fmt.Println("error opening browser:", err)
		return
	}

}
