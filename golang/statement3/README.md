# Trie CLI in Go

This program implements a simple command-line interface (CLI) in Go that allows users to manage a collection of strings using a trie data structure.

The application runs in a loop, accepts user commands, and performs operations such as adding, listing, checking, and removing strings.

---

## Features

The program supports the following commands:

- **add** – Add a string to the store  
- **list** – List all stored strings (DFS is used to traverse the trie)  
- **check** – Check if a string exists  
- **remove** – Remove a string from the store  
- **quit** – Exit the program  

---

## How to Run

```bash
go run trie.go
go test
