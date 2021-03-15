package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	file, _ := os.OpenFile("./test.txt", os.O_RDWR, os.ModePerm)
	for i := 0; i < 1_000_000; i++ {
		_, err := io.WriteString(file, randomCharacters()+"\n")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func randomCharacters() (s string) {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	for i := 0; i < 5; i++ {
		s += string(chars[rand.Intn(len(chars)-1)])
	}

	return s
}
