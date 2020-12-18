package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	dict := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	line := ""
	keyPair := make([]string, 2)

	// Loop and read input
	for scanner.Scan() {

		line = scanner.Text()

		// End loop on a blank line
		if line == "" {
			break
		}

		// Split line
		keyPair = strings.Split(line, " ")

		// Insert key and value into dict
		dict[keyPair[1]] = keyPair[0]
	}

	var translation strings.Builder

	// Translate by looking up words
	for scanner.Scan() {

		line = scanner.Text()

		// End loop on a blank line
		if line == "" {
			break
		}

		// Look for word in dictionary
		if val, exists := dict[line]; exists {
			translation.WriteString(val + "\n")
		} else {
			translation.WriteString("eh\n")
		}
	}

	fmt.Print(translation.String())
}
