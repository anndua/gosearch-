# s — Terminal Search

A small Go CLI for quickly opening web searches from the terminal.

Instead of opening a browser, typing a query, and choosing a site manually:

```bash
s golang
```

opens the search directly in Chrome.

## Features

- Google search by default
- GitHub search
- YouTube search
- Go documentation search
- Chrome browser support
- Multi-word queries
- Colored terminal output
- Custom CLI help

## Usage

```bash
s <query>
```

Google is the default search engine:

```bash
s golang
s "golang concurrency"
s linux networking
```

### Search GitHub

```bash
s -gh golang
s -gh "go cli"
```

### Search YouTube

```bash
s -yt "golang tutorial"
s -yt linux networking
```

### Search Go documentation

```bash
s -docs "net/http"
s -docs goroutines
```

### Browser

Chrome is currently the default:

```bash
s -browser chrome golang
```

## Help

```bash
s -h
```

Example output:

```text
S - Terminal Search

Usage:
  s [options] <query>

Search Engines:
  -google
      Search Google (default)
  -gh
      Search GitHub
  -yt
      Search YouTube
  -docs
      Search Go documentation

Options:
  -browser string
      browser to open (default: chrome)
  -h, -help
      show this help
```

## Installation

### Requirements

- Linux / WSL
- Go
- Google Chrome installed on Windows when running under WSL

### Build

Clone the repository and enter the project:

```bash
git clone <YOUR_REPOSITORY_URL>
cd gosearch
```

Build the binary:

```bash
go build -o gosearch
```

Install it into your local executable directory:

```bash
mkdir -p ~/.local/bin
cp gosearch ~/.local/bin/gosearch
```

Add `~/.local/bin` to your `PATH` if it is not already there:

```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

You can then use:

```bash
gosearch golang
```

If you want the short command:

```bash
ln -s ~/.local/bin/gosearch ~/.local/bin/s
```

Then:

```bash
s golang
```

## Development

Run directly from the source:

```bash
go run . golang
```

Format the code:

```bash
gofmt -w main.go
```

Build:

```bash
go build -o gosearch
```

Check the project:

```bash
go vet ./...
```

## Project Structure

```text
gosearch/
├── .gitignore
├── go.mod
├── main.go
└── README.md
```

The project currently uses Go's standard library only.

## Roadmap

This project is intentionally being built incrementally while learning Go.

- [x] Basic CLI
- [x] Read terminal arguments
- [x] Build search URLs
- [x] Open Chrome from WSL
- [x] Google search
- [x] GitHub search
- [x] YouTube search
- [x] Go documentation search
- [x] Colored terminal output
- [ ] Refactor into functions
- [ ] Better argument parsing
- [ ] More browser support
- [ ] Search history
- [ ] Configuration file
- [ ] Custom aliases
- [ ] Tests
- [ ] Cross-platform support
- [ ] GitHub releases

## Why?

This started as a small project for learning Go through something practical.

The goal is to turn it into a useful standalone CLI while learning:

- Go
- Linux
- WSL
- CLI design
- HTTP and URLs
- Git and GitHub
- testing
- process management
- configuration
- eventually concurrency and more advanced Go concepts

## License

MIT
