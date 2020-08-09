package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(c rune) rune {
	mapping := make(map[rune]rune)
	mapping['A'] = 'N'
	mapping['B'] = 'O'
	mapping['C'] = 'P'
	mapping['D'] = 'Q'
	mapping['E'] = 'R'
	mapping['F'] = 'S'
	mapping['G'] = 'T'
	mapping['H'] = 'U'
	mapping['I'] = 'V'
	mapping['J'] = 'W'
	mapping['K'] = 'X'
	mapping['L'] = 'Y'
	mapping['M'] = 'Z'
	mapping['N'] = 'A'
	mapping['O'] = 'B'
	mapping['P'] = 'C'
	mapping['Q'] = 'D'
	mapping['R'] = 'E'
	mapping['S'] = 'F'
	mapping['T'] = 'G'
	mapping['U'] = 'H'
	mapping['V'] = 'I'
	mapping['W'] = 'J'
	mapping['X'] = 'K'
	mapping['Y'] = 'L'
	mapping['Z'] = 'M'

	lowCaseMapping := make(map[rune]rune)
	for k, v := range mapping {
		lowCaseMapping[unicode.ToLower(k)] = unicode.ToLower(v)
	}
	for k, v := range lowCaseMapping {
		mapping[k] = v
	}
	value := mapping[c]
	if value == 0 {
		return c
	} else {
		return value
	}
}

func (r *rot13Reader) Read(bytes []byte) (n int, e error) {
	input := make([]byte, 1024)
	readCount, err := r.r.Read(input)
	if err != nil {
		return 0, err
	}

	for i := 0; i < readCount; i++ {
		bytes[i] = byte(rot13(rune(input[i])))
	}

	return len(bytes), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
