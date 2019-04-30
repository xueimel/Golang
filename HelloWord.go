package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*HelloWord.go is a implementation of the golang based off of a finite automata.
An input file is provided as an argument and a string is taken from the user during
execution. The string is then broken to characters which are compared to the input
file. If the character matches, the string moves to the next character, else it
continues to move forward until a match is found or end of file is reached.*/
type Word struct {
	word   []byte
	index  int
	length int
}

//methods for struct
func (w *Word) getChar() byte { //return the present char
	x := w.index
	return w.word[x]
}
func (w *Word) reachedEnd() bool { //return if all chars have been traversed
	if w.index+1 == w.length {
		return true
	}
	return false
}
func (w *Word) updateIndex() { w.index++ } //move to next char

func main() {
	traversed := false     //if the entire input string has been searched
	fileName := os.Args[1] // input file to be traversed
	file, err := ioutil.ReadFile(fileName)
	if os.IsNotExist(err) { //catches file not found
		fmt.Println(fileName, " does not exist")
		os.Exit(1)
	}
	fileStr := string(file) //converts file from bytes to string
	fmt.Printf("Enter the string of characters you wish to search for:")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')                          //read from command line
	str = strings.TrimSuffix(str, "\n")                        //remove trailing newline
	str = strings.TrimSuffix(str, "\r")                        // remove carriage for windows
	words := []byte(str)                                       //convert input string to bytes for comparison
	keyWord := Word{word: words, index: 0, length: len(words)} //make struct
	for x, _ := range fileStr {
		if !keyWord.reachedEnd() { //if all characters of string have been traversed
			if keyWord.getChar() == byte(fileStr[x]) {
				keyWord.updateIndex() //move to next character
			}
		} else {
			traversed = true //all chars traversed
			break
		}
	}
	if traversed {
		fmt.Printf("The Characters in the string, %s were found throughout the file %s", str, fileName)
	} else {
		fmt.Printf("The file, %s did not contain the characters, %s you provided in the appropriate order to complete the match", fileName, str)
	}
}
