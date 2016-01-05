package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

const (
	unknownFormat uint = iota
	jsonFormat
	xmlFormat
)

func main() {
	pretty(os.Stdin, os.Stdout)
}

func pretty(r io.Reader, w io.Writer) (int64, error) {

	buf := bufio.NewReaderSize(r, 4)

	var format uint
	for {
		ch, _, err := buf.ReadRune()
		if err != nil {
			return 0, err
		}
		if f, ok := formats(ch); ok {
			format = f
			if err := buf.UnreadRune(); err != nil {
				return 0, err
			}
			break
		}
		if endOfLine(ch) {
			return 0, errors.New("unable to recognize this format")
		}
	}

	b, err := ioutil.ReadAll(buf)
	if err != nil {
		return 0, err
	}

	var out bytes.Buffer
	switch format {
	case jsonFormat:
		if err := json.Indent(&out, b, "", "\t"); err != nil {
			return 0, err
		}
	case xmlFormat:
		return 0, errors.New("xml is not yet supported")
	default:
		return 0, errors.New("known format error, please file a bug")
	}

	return out.WriteTo(w)
}

func endOfLine(ch rune) bool {
	return ch == '\n' || ch == '\r'
}

func formats(ch rune) (uint, bool) {
	switch ch {
	case '{':
		return jsonFormat, true
	case '<':
		return xmlFormat, true
	default:
		return unknownFormat, false
	}
}
