package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"prefixMatch/pkg/files"
	"prefixMatch/pkg/matcher"
	"sort"
)

func main() {
	fmt.Println("Enter the file path: ")
	var path string
	_, err := fmt.Scanln(&path)
	if err != nil {
		return
	}
	fmt.Println("Enter the search string: ")
	var word string
	_, err = fmt.Scanln(&word)
	if err != nil {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var prefixWords []string
	r := bufio.NewReader(file)
	s, e := files.ReadLine(r)
	for e == nil {
		prefixWords = append(prefixWords, s) // read the prefixes from file add to an array
		s, e = files.ReadLine(r)
	}
	sort.Strings(prefixWords) // sort the array of prefixes, as preprocessing is allowed.
	res := matcher.LongestPrefix(prefixWords, word)
	fmt.Println(res)
}
