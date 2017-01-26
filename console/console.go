package console

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// AskYN asks in stdout from stdin to type y||Y||n||N||[ret]
// Panic when bufio reader returns error when reading a string
func AskYN(msg string) bool {
	return askYN(msg, os.Stdout, os.Stdin)
}

// askYN asks in w from rd reading y||Y||n||N||[ret]
// Panic when bufio reader returns error when reading a string
func askYN(msg string, w io.Writer, rd io.Reader) bool {
	var in string
	var err error
	msg += " (y/N) "
	reader := bufio.NewReader(rd)
	fmt.Fprint(w, msg)

	for {
		in, err = reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		if in == "\n" || in == "y\n" || in == "Y\n" || in == "n\n" || in == "N\n" {
			break
		}

		fmt.Fprint(w, msg)
	}

	return in == "y\n" || in == "Y\n"
}
