# Word Link- Finding connections between five letter words

Given any two five letter words, this program will find the shortest path between them using intermediate words that are only one letter different.

For example, starting at "start" and ending at "close" would yield the path start -> stare -> share -> shore -> chore -> chose -> close.

Certain words cannot be linked in this way. For example, starting at "zebra" and ending at "horse" yields no path because no path exists between the two words.

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
