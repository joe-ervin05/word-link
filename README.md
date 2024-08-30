# Word Link- Connecting Five Letter Words

Given any two five letter words, this program will find the shortest path between them using intermediate words that are only one letter different.

For example, starting at "heads" and ending at "tails" would yield the path: heads -> heals -> hells -> halls -> hails -> tails.

Certain words cannot be linked in this way. For example, starting at "zebra" and ending at "horse" yields no path because no path exists between the words.

## Implementation

The program uses a word bank of 5757 five letter English words. It uses this word bank to create a graph that holds each word (the verticies) and all words that are only one letter different from each word (the edges).

When a start and end word are given, the program performs a breadth-first search from the start word to all other connected words. It then uses the result of this search to find the shortest path between the start word and the end word.

Breadth-first search is used for this problem because it has the property that the first time it reaches a vertex is the shortest path to that vertex in an unweighted graph.

Many paths between words are cyclical which causes some implementations of breadth-first search to get stuck in an infinite loop. The solution to this is to keep track of which words have already been visited and only visit a new word if it has not been visited yet.

## Running the program

To run the program, you must have go installed.

to run with make, call:

```
    make run
```

otherwise, run something like:

```
    go build -o ./tmp/main *.go && ./tmp/main
```
