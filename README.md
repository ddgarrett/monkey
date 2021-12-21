# Monkey

From ["Writing An Interpreter In Go"](https://interpreterbook.com/). 

This implements an interpreter for the [Monkey programming language](https://monkeylang.org/) as defined by Thorsten Ball.

Why do this? A few reasons:

1. I've always been curious about interpreters and transpilers (source to source compilers) since my first job involved using a transpiler.

2. Once I have an implementation of an interpreter using an abstract syntax tree (AST) it might be fun
to extend it to add new programming elements, creating my own scripting language.

3. One of the extensions that might be fun to add is an implementation of the [MonGolang](https://github.com/ddgarrett/mongolang) library, an "easier" to use implementation of MongoDB access in Go, which might in turn provide a scripting language for using MongoDB.

4. I want to continue to learn more about the Go programming language itself and best practices in programming with Go.

#### December 1, 2021 Update

There are now two new branches, <b>interpreter</b> and <b>compiler</b>, both based on the master branch. The interpreter branch contains code from ["Writing An Interpreter In Go"](https://interpreterbook.com/). The compiler branch contains code from ["Writing A Compiler In Go"](https://compilerbook.com/).

Long term, the intent is to keep the master branch an interpreter based version of the Monkey language. Accordingly, changes may be merged from the interpreter branch into the master branch, and then into the compiler branch. However it is not anticipated that changes to the compiler branch will be merged into the master branch.

#### December 21, 2021 Update

In ["Writing A Compiler In Go"](https://compilerbook.com/) code from 
["Writing An Interpreter In Go"](https://interpreterbook.com/) is refactored to move  built-in function logic from the `evaluator` package to the `object` package. This allows the interpreter and compiler to use the same codebase for built-in functions. This update was made first to the interpreter branch, used to update the master branch and then pulled into the compiler branch.