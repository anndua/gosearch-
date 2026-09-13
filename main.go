package main

import (
	"flag"
	"fmt"
	"net/url"
	"os/exec"
	"strings"
)

func main() {
	browser := flag.String("browser", "chrome", "browser to open")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Usage: gosearch [engine] <query>")
		return
	}

	args := flag.Args()

	engine := "google"
	queryStart := 0

	if args[0] == "gh" {
		engine = "github"
		queryStart = 1
	} else if args[0] == "yt" {
		engine = "youtube"
		queryStart = 1
	} else if args[0] == "docs" {
		engine = "docs"
		queryStart = 1
	}

	if len(args) <= queryStart {
		fmt.Println("Missing search query")
		return
	}

	query := strings.Join(args[queryStart:], " ")

	var searchURL string

	switch engine {
	case "google":
		searchURL = "https://www.google.com/search?" +
			url.Values{"q": []string{query}}.Encode()

	case "github":
		searchURL = "https://github.com/search?" +
			url.Values{"q": []string{query}}.Encode()

	case "youtube":
		searchURL = "https://www.youtube.com/results?" +
			url.Values{"search_query": []string{query}}.Encode()

	case "docs":
		searchURL = "https://pkg.go.dev/search?" +
			url.Values{"q": []string{query}}.Encode()
	}

	fmt.Println("Searching", engine, "for:", query)

	var cmd *exec.Cmd

	switch *browser {
	case "chrome":
		cmd = exec.Command(
			"/mnt/c/Program Files/Google/Chrome/Application/chrome.exe",
			searchURL,
		)

	default:
		fmt.Println("Unknown browser:", *browser)
		return
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("error opening browser:", err)
		return
	}
}
