package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sb := strings.Builder{}
	body := false // False until we find something we assume is the JSON body
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			// If the line begins with "{" or "[", we guess is the start JSON body
			if line[:1] == "{" || line[:1] == "[" {
				body = true
			}
		}
		if body {
			// JSON body is written into a string builder to be later processed
			sb.WriteString(line)
			sb.WriteString("\n")
		} else {
			// Anything before the body is printed out as is.
			fmt.Println(line)
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
			fmt.Println("---")
			fmt.Print(err)
			os.Exit(1)
		}
		fmt.Println(out.String())
	}
}
