package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMd5Str(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	s := "hello"
	m := GetMd5Str(s)
	fmt.Printf("%s md5 is %s\n", s, m)
}
