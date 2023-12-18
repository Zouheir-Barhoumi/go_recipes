package bytestrings

import (
	"bytes"
	"io"
)

func Buffer(rawString string) *bytes.Buffer {
	// We'll start with a string encoded into raw bytes
	rawBytes := []byte(rawString)
	// there are mutliple ways to create a buffer form the raw bytes
	// or form the original string
	var b = new(bytes.Buffer)

	b.Write(rawBytes)
	// alternatively
	// b = bytes.NewBuffer(rawBytes)
	// and avoiding the initial byte array altogether
	// b = bytes.NewBufferString(rawString)
	return b
}

func toString(r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
