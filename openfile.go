package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

// Opens a file, and returns its size in bytes, and a bytes buffer of its contents.
func openFile(name string) (byteCount int, buffer *bytes.Buffer) {

	data, err = os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	reader := bufio.NewReader(data)
	buffer = bytes.NewBuffer(make([]byte, 0))
	part = make([]byte, chunksize)

	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if err != io.EOF {
		log.Fatal("Error Reading ", name, ": ", err)
	} else {
		err = nil
	}

	byteCount = buffer.Len()
	return
}
