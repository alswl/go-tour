package main

import "golang.org/x/tour/reader"

type MyReader struct {}

const charA = int('A')

func (r MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = byte(charA)
	}
	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})

}
