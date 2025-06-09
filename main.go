package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const endLine = '\n'

func main() {
	reader := bufio.NewReader(os.Stdin)
	sb := strings.Builder{}
	for {
		body := false // False until we find something we assume is the JSON body
		line, err := reader.ReadString(endLine)

		if line != "" {
			// If the line begins with "{" or "[", we guess is the start JSON body
			if strings.HasPrefix(strings.TrimSpace(line), "{") || strings.HasPrefix(strings.TrimSpace(line), "[") {
				body = true
			}
		}
		if body {
			// JSON body is written into a string builder (without end-line) to be later processed
			sb.WriteString(strings.TrimSpace(line))
		} else {
			// Anything before that is not JSON body is printed out as is.
			fmt.Print(line)
		}
		if err != nil {
			if errors.Is(err, io.EOF) { // No more lines
				break
			} else {
				errorExit(err)
			}
		}
	}

	var out bytes.Buffer
	payload := sb.String()
	if len(payload) > 0 {
		// Prettify the assumed JSON body
		err := json.Indent(&out, []byte(payload), "", "  ")
		if err != nil {
			// Maybe it cannot be parsed, maybe something else... Print it out as we read it
			fmt.Println(payload)
			// Print out the error at the end.
			errorExit(err)
		}
		fmt.Println(out.String())
	}
}

func errorExit(err error) {
	fmt.Println("---")
	fmt.Print(err)
	os.Exit(1)
}
