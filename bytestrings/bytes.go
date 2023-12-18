package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

func WorkWithBuffer() error {
	rawString := "It's easy to encode unicode into a byte array"

	b := Buffer(rawString)

	// We can quickly convert a buffer back into bytes with
	// b.Bytes() or a string with b.String()
	fmt.Println(b.String())

	// because this is an io Reader we can make use of
	// generic io reader functions such as
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	// we can also take out bytes and careate a bytes reader
	// these readers implement io.Reader, io.ReaderAt,
	// io.WriterTo, io.Seeker, io.ByteScanner, and
	// io.RuneScanner interfaces

	reader := bytes.NewReader([]byte(rawString))

	// we can also plug it into a scanner that allows
	// buffered reading and tokenization
	// scanner, _, err := bufio.ScanWords(reader) // TODO: fix this (deprecated)
	// scanner.Split(bufio.ScanWords)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}
