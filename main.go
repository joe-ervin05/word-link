package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {

	wrdList, err := readLines("words.txt")
	if err != nil {
		log.Fatal(err)
	}

	wrds, err := loadGraph()

	if err != nil {
		err = _initGraph()

		if err != nil {
			log.Fatal(err)
		}

		wrds, err = loadGraph()

		if err != nil {
			log.Fatal(err)
		}
	}

	idx1, idx2 := getInp(wrdList)

	for idx1 == idx2 {
		fmt.Println("cannot link a word to itself")

		idx1, idx2 = getInp(wrdList)
	}

	path := wrds.shortestPath(idx1, idx2)

	if path == nil {
		fmt.Printf("%s and %s cannot be linked\n", wrdList[idx1], wrdList[idx2])
		return
	}

	for i := len(path) - 1; i >= 0; i-- {
		fmt.Printf(" %s", wrdList[path[i]])
		if i > 0 {
			fmt.Printf(" ->")
		}
	}

}

func getInp(wrdList []string) (idx1 int, idx2 int) {

	var wrd1 string

	fmt.Print("start word: ")
	fmt.Scan(&wrd1)

	idx1 = indexOf(wrdList, strings.ToLower(wrd1))

	for idx1 < 0 {
		if len(wrd1) != 5 {
			fmt.Print("your word must be five letters")
		} else {
			fmt.Print("word not found")
		}

		fmt.Print("\nstart word: ")

		fmt.Scan(&wrd1)
		idx1 = indexOf(wrdList, strings.ToLower(wrd1))
	}

	var wrd2 string

	fmt.Print("end word: ")
	fmt.Scan(&wrd2)

	idx2 = indexOf(wrdList, strings.ToLower(wrd2))

	for idx2 < 0 {
		if len(wrd2) != 5 {
			fmt.Print("your word must be five letters")
		} else {
			fmt.Print("word not found")
		}

		fmt.Print("\nend word: ")

		fmt.Scan(&wrd2)
		idx2 = indexOf(wrdList, strings.ToLower(wrd2))
	}

	return idx1, idx2
}
