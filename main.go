package main

import (
	"bytes"
	"compress/zlib"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	var input string
	flag.StringVar(&input, "f", "", "input bin file")
	flag.Parse()

	org := readfile(input)
	l := len(org)
	b := make([]byte, l)
	copy(b, org)
	start := time.Now()
	fixfound := 0

	for i := 2; i <= l; i++ {
		ob := b[i]
		for j := 0; j <= 255; j++ {
			b[i] = byte(j)
			if b[i] == ob {
				continue
			}
			op, err := uncompress(b, true)
			if err == nil {
				fmt.Printf("OK: fixed when b[%d] change to %x \n", i, j)
				fixfound++
				ioutil.WriteFile(fmt.Sprintf("./%s.fix.%d.bin", input, fixfound), op, 0644)
			}
		}
		b[i] = ob

		t := time.Now()
		elapsed := t.Sub(start)
		eta := time.Duration(elapsed.Nanoseconds() * int64(l-i) / int64(i))
		elapsed = elapsed.Round(time.Second)
		eta = eta.Round(time.Second)
		fmt.Println("trying byte:", i, "/", l, " took:", elapsed, " ETA:", eta, " fix found:", fixfound)
	}
}

func readfile(f string) []byte {
	dat, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalln("ReadFile Error", err)
	}
	return dat
}

func uncompress(buff []byte, slient bool) ([]byte, error) {
	b := bytes.NewReader(buff)

	r, err := zlib.NewReader(b)
	if err != nil {
		log.Fatalln("zlib NewReader error", err)
	}

	defer r.Close()

	buff, e := ioutil.ReadAll(r)
	if e != nil && !slient {
		log.Println("uncompress error", e)
	}

	return buff, e
}

func prints(b []byte) {
	for i := 0; i < 12; i++ {
		fmt.Printf("%x ", b[i])
	}
	fmt.Print("...")

	l := len(b)
	for i := l - 10; i < l; i++ {
		fmt.Printf("%x ", b[i])
	}
	fmt.Print("\n")
}
