package main

import (
	"debug/pe"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: dumpe <BINARY>")
		os.Exit(0)
	}
	path := os.Args[1]
	file, err := pe.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	text := file.Section(".text")
	offset := text.SectionHeader.Size
	size := text.SectionHeader.VirtualSize

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buffer, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	for _, v := range buffer[offset : offset+size] {
		fmt.Printf("\\x%02x", v)
	}
}
