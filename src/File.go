package main

import (
	"fmt"
	// "io"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkSum(fname string, start int, size int) []string {
	f, err := os.Open(fname)
	checkError(err)
	defer f.Close()
	checkSums := []string{}
	b := make([]byte, size)
	offset := 0
	for {
		f.Seek(offset, 0)
		n, err := f.Read(b)
		checkError(err)
		if n == 0 {
			return checkSums
		}
		offset += size
		append(checkSums, GetMd5Str(string(b)))
	}

}

func main() {
	fmt.Println(checkSum("/tmp/dat", 0, 1))
}
