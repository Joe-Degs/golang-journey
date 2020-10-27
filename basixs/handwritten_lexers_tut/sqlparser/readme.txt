Learning about handwritten parsers and lexers in go
---------------------------------------------------

Its actually kinda cool, the way the go language is structured and 
well suited for this kind of stuff. I wrote a piece of code for an
exercism exercise recently and it felt good the way go gave you 
absolute control over what your program does. The abstraction is minimal and 
concise and leaves the programmer to think about code the way it should be.

I also learnt a ton of stuff from dealing with runes and ascii char codes. It felt
good to dive into stuff like that, it felt badass that i was actually comparing
runes of different characters and doing stuff with them and not calling some
builtin packages to do it for me.

I strolled around the internet answers to all questions i had writing the code and
stumbled upon a blog post on gopher academy on handwritten lexers and parsers in go.
I felt like why not try to go through step by step and see what you learn at the end.
So here i go trying to follow the blog post and put down my thoughts in the process.

DRUM ROLL, LETS GET STARTED...

First of what the fuck is a lexer and a parser?. To know more you could read up on
lexical scannning and parsing on wikipedia.
A lexer breaks a series of input into tokens and feeds them into a parse to make sense
out of it. To parse something like formal english grammar, we first break them into
series of characters(tokens) that have a unit of meaning to us and then feed those
tokens to a parser to make meaningful sense out of a group of tokens.

Ex. For an sql-like language, the statement:
SELECT * FROM mytable

Can be broken down into or tokenized to;
`SELECT` `WS` `ASTERISK` `WS` `FROM` `WS` `STRING<mytable>`

This process is called lexical scanning and its similar to how we break down words
in formal language when we read.

From what i'm reading on the internet, there are lots of tools for lexer and parser 
generation around. But i guess writing it in go would be so much more fun to do.

This tutorial is about writing a lexer and a parser for sql select statements.
Time to write some go code now :)

Done with the coding and it was a fun and nice tutorial, also learnt tonnes of stuff from it.
Parsing once you start getting into it is not as difficult as it seems once you can see
the bigger picture. Planning should always be the first thing to do when looking to build 
software. lets always resort to planning as an initial stage of the software design process
seeing the bigger picture helps to write better and beautiful code. JOE, always remember to put
enough thought into the structure and reason of the program before you code. Understand the problem
statement and try to have a clear mental picture of how your program to work before you start to code.
Don't be afraid to change along the road too. If you wrote it well you can write it better.
