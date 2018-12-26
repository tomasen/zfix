package main

import (
	"fmt"
	"hash/adler32"
	"testing"
)

func TestOpen(t *testing.T) {
	b := readfile("ok.bin")
	prints(b)
	buff, err := uncompress(b, false)
	fmt.Println(err)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Checksum %x Length: %d\n", adler32.Checksum(buff), len(buff))

	b = readfile("fail.bin")
	prints(b)

	buff, err = uncompress(b, false)
	fmt.Println(err)
	if err == nil {
		t.Fatal(err)
	}

	fmt.Printf("Checksum %x Length: %d\n", adler32.Checksum(buff), len(buff))
	prints(buff)
}
