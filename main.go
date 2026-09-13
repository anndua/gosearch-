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

	flag.Usage = func() {
		fmt.Println()
		fmt.Println(bold + cyan + "  S - Terminal Search" + reset)
		fmt.Println()
		fmt.Println("  Usage:")
		fmt.Println("    s [options] <query>")
		fmt.Println()
		fmt.Println("  Options:")
		fmt.Println("    -browser string")
		fmt.Println("        browser to open (default: chrome)")
		fmt.Println("    -h, -help")
		fmt.Println("        show this help")
		fmt.Println()
		fmt.Println("  Examples:")
		fmt.Println("    s golang")
		fmt.Println("    s \"golang concurrency\"")
		fmt.Println("    s -browser chrome linux")
		fmt.Println()
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	query := strings.Join(flag.Args(), " ")

	params := url.Values{}
	params.Set("q", query)

	searchURL := "https://www.google.com/search?" + params.Encode()

	fmt.Println()
	fmt.Println(cyan + "  Searching Google for:" + reset)
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
