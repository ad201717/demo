package main

import (
	// "bufio"
	// "fmt"
	// "io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("/tmp/dat2")
	check(err)
	f.Close()

	err = os.Rename("/tmp/dat2", "/tmp/dat3")
	check(err)
}
