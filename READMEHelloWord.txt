HelloWord.go is a implementation of the golang based off of a finite automata.
An input file is provided as an argument and a string is taken from the user during
execution. The string is then broken to characters which are compared to the input
file. If the character matches, the string moves to the next character, else it
continues to move forward until a match is found or end of file is reached.

to compile and run:
	go run HelloWord.go <input file>
