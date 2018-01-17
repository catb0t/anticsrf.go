package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func microtime() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

func random_key(length byte) string {
	bytes := make([]byte, length/2)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(random_key(byte(i)))
	}
	fmt.Println(microtime())
}
