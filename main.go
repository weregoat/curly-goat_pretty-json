package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer

	// Find the first '{' that we assume mark the beginning of the JSON payload
	// Big assumption in general, not so big in the context this script was
	// written for.
	i := bytes.IndexByte(b, byte('{'))
	if i > 0 {
		fluff := b[:i]
		payload := b[i:]
		fmt.Fprintf(os.Stdout, "%s", fluff)
		err = json.Indent(&out, payload, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		out.WriteTo(os.Stdout)
		fmt.Fprintln(os.Stdout, "")
	} else {
		fmt.Fprintf(os.Stdout, "%s", b)
	}

}
