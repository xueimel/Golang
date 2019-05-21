HelloWord.go is a implementation of the golang based off of a finite automaton or regex
search.
User provides an input file to be parsed before execution and a target string to 
be searched for during execution.
The string is then broken to characters which are compared to sequential characters
in the input file. If a character is found to match, the target string moves to the 
next character in the sequence, else it continues to move forward until a match 
is found or end of file is reached.

to compile and run:
	go run HelloWord.go <input file>
