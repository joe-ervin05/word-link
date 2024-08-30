package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"math"
	"os"
)

const WORD_COUNT = 5757

type queue []int
type wordGraph [WORD_COUNT][]int

// takes in a queue and returns the top item
func top(q queue) int {
	return q[0]
}

func discard(q queue) queue {
	return q[1:]
}

func isEmpty(q queue) bool {
	return len(q) == 0
}

// preforms a breadth-first search starting at the first word index
// returns arrays for each words parent and their distance to the starting node
func (wg wordGraph) bfs(from int) (par, dist []int) {
	q := queue{from}

	dist = make([]int, WORD_COUNT)
	par = make([]int, WORD_COUNT)

	for i := 0; i < WORD_COUNT; i++ {
		par[i] = -1
		dist[i] = math.MaxInt32
	}

	dist[from] = 0

	q = append(q, from)

	for !isEmpty(q) {

		node := top(q)
		q = discard(q)

		for _, wrd := range wg[node] {
			if dist[wrd] == math.MaxInt32 {
				par[wrd] = node
				dist[wrd] = dist[node] + 1

				q = append(q, wrd)
			}
		}
	}

	return par, dist

}

// takes in a start word index and destination word index
// returns an array with the indexes for the shortest path between the two words
func (wg wordGraph) shortestPath(s, d int) []int {
	par, dist := wg.bfs(s)

	if dist[d] == math.MaxInt32 {
		return nil
	}

	path := []int{d}

	currNode := d

	for par[currNode] != -1 {
		path = append(path, par[currNode])
		currNode = par[currNode]
	}

	return path
}

// returns a graph of word indexes and the words they are linked to
func loadGraph() (wrdGraph wordGraph, err error) {
	data, err := os.ReadFile("wordgraph.gob")

	if err != nil {
		return wrdGraph, err
	}

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	err = dec.Decode(&wrdGraph)

	return wrdGraph, err
}

func _initGraph() error {
	wrds, err := readLines("words.txt")

	if err != nil {
		return err
	}

	wrdGraph := wordGraph{}

	for i := 0; i < len(wrds); i++ {
		for j := 0; j < len(wrds); j++ {

			if misses(wrds[i], wrds[j]) == 1 {
				wrdGraph[i] = append(wrdGraph[i], j)
			}

		}

	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err = enc.Encode(wrdGraph)
	if err != nil {
		return err
	}

	return os.WriteFile("wordgraph.gob", buf.Bytes(), 0644)

}

func misses(wrd1, wrd2 string) (m int) {
	for i := 0; i < len(wrd1); i++ {
		if wrd1[i] != wrd2[i] {
			m++
		}
	}

	return m
}

func readLines(path string) (lines []string, err error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scn := bufio.NewScanner(file)

	for scn.Scan() {
		lines = append(lines, scn.Text())
	}

	if scn.Err() != nil {
		return nil, err
	}

	return lines, nil
}

func indexOf(list []string, str string) int {

	low := 0
	high := WORD_COUNT - 1

	for low <= high {
		mid := low + (high-low)/2

		if list[mid] == str {
			return mid
		}

		if list[mid] < str {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
