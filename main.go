package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)

const (
	unknownFormat uint = iota
	jsonFormat
	xmlFormat
)

func main() {
	if err := pretty(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func pretty(r io.Reader, w io.Writer) error {

	buf := bufio.NewReaderSize(r, 4)

	var format uint
	for {
		ch, _, err := buf.ReadRune()
		if err != nil {
			return err
		}

		if f, ok := formats(ch); ok {
			format = f
			if err := buf.UnreadRune(); err != nil {
				return err
			}
			break
		}

		if endOfLine(ch) {
			return errors.New("unable to recognize this format")
		}
	}

	b, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}

	switch format {
	case jsonFormat:
		var out *bytes.Buffer
		if err := json.Indent(out, b, "", "\t"); err != nil {
			return err
		}
		if _, err := out.WriteTo(w); err != nil {
			return err
		}
	case xmlFormat:
		d := xml.NewDecoder(bytes.NewReader(b))
		e := xml.NewEncoder(w)
		e.Indent("", "\t")

		for {
			t, err := d.Token()
			if err == io.EOF {
				break
			}
			if tok, ok := t.(xml.CharData); ok {
				r, _ := utf8.DecodeRune(tok)
				if blank(r) {
					continue
				}
			}
			e.EncodeToken(t)
		}
		return e.Flush()
	default:
		return errors.New("known format error, please file a bug")
	}
	return nil
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

func blank(ch rune) bool {
	return ch == ' ' || ch == '\n' || ch == '\t'
}
