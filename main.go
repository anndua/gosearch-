package main

import (
	"flag"
	"fmt"
	"net/url"
	"os/exec"
	"strings"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	cyan   = "\033[36m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	bold   = "\033[1m"
)

func main() {
	browser := flag.String("browser", "chrome", "browser to open")

	google := flag.Bool("google", false, "search Google")
	github := flag.Bool("gh", false, "search GitHub")
	youtube := flag.Bool("yt", false, "search YouTube")
	docs := flag.Bool("docs", false, "search Go documentation")

	flag.Usage = func() {
		fmt.Println()
		fmt.Println(bold + cyan + "  S - Terminal Search" + reset)
		fmt.Println()
		fmt.Println("  Usage:")
		fmt.Println("    s [options] <query>")
		fmt.Println()
		fmt.Println("  Search Engines:")
		fmt.Println("    -google")
		fmt.Println("        Search Google (default)")
		fmt.Println("    -gh")
		fmt.Println("        Search GitHub")
		fmt.Println("    -yt")
		fmt.Println("        Search YouTube")
		fmt.Println("    -docs")
		fmt.Println("        Search Go documentation")
		fmt.Println()
		fmt.Println("  Options:")
		fmt.Println("    -browser string")
		fmt.Println("        browser to open (default: chrome)")
		fmt.Println("    -h, -help")
		fmt.Println("        show this help")
		fmt.Println()
		fmt.Println("  Examples:")
		fmt.Println("    s golang")
		fmt.Println("    s -gh golang")
		fmt.Println("    s -yt \"golang tutorial\"")
		fmt.Println("    s -docs \"net/http\"")
		fmt.Println()
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	// Make sure only one search engine was selected.
	engineCount := 0

	if *google {
		engineCount++
	}

	if *github {
		engineCount++
	}

	if *youtube {
		engineCount++
	}

	if *docs {
		engineCount++
	}

	if engineCount > 1 {
		fmt.Println(red + "  Error:" + reset + " choose only one search engine")
		return
	}

	query := strings.Join(flag.Args(), " ")

	var searchURL string
	var engineName string

	switch {
	case *github:
		engineName = "GitHub"
		searchURL = "https://github.com/search?" +
			url.Values{"q": []string{query}}.Encode()

	case *youtube:
		engineName = "YouTube"
		searchURL = "https://www.youtube.com/results?" +
			url.Values{"search_query": []string{query}}.Encode()

	case *docs:
		engineName = "Go Docs"
		searchURL = "https://pkg.go.dev/search?" +
			url.Values{"q": []string{query}}.Encode()

	default:
		engineName = "Google"
		searchURL = "https://www.google.com/search?" +
			url.Values{"q": []string{query}}.Encode()
	}

	fmt.Println()
	fmt.Println(cyan + "  Searching " + engineName + " for:" + reset)
	fmt.Println("  " + bold + query + reset)
	fmt.Println()

	var cmd *exec.Cmd

	switch *browser {
	case "chrome":
		cmd = exec.Command(
			"/mnt/c/Program Files/Google/Chrome/Application/chrome.exe",
			searchURL,
		)

	default:
		fmt.Println(red+"  Error:"+reset+" unsupported browser:", *browser)
		return
	}

	fmt.Println(green + "  ✓ Opening Chrome..." + reset)

	err := cmd.Run()
	if err != nil {
		fmt.Println(red+"  ✗ Error opening browser:"+reset, err)
		return
	}

	fmt.Println(green + "  ✓ Done" + reset)
}
