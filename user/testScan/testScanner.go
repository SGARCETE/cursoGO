package main

import (
	"github.com/user/Scanner"
	"strings"
	"fmt"
	"os"
)

func main() {
	const input = "hola, como ,va,"
	scanner := Scanner.NewScanner(strings.NewReader(input))
	scanner.Split(Scanner.SeparatorSplitter{','})

	for scanner.Scan() {
		fmt.Printf("%q", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
