package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestKemono(t *testing.T) {
	file, err := os.Open("samplecode\\helloworld.kemono")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		input, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		input = replaceKemono(input)
		fmt.Println(string(input))
	}
}
