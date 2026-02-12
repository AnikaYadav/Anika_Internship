package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

func main() {
	items := make(map[string]bool)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Command (add, list, remove, check, quit):")
		fmt.Print("> ")

		command, err := readInput(reader)
		if err != nil {
			fmt.Println("Error reading command:", err)
			continue
		}
		command = strings.TrimSpace(strings.ToLower(command))

		switch command {

		case "add":
			fmt.Print("Input: ")
			input, err := readInput(reader)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			addItem(items, input)
			fmt.Println("Added!")

		case "check":
			fmt.Print("Input: ")
			input, err := readInput(reader)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			if checkItem(items, input) {
				fmt.Println("Exists!")
			} else {
				fmt.Println("Does not exist.")
			}

		case "remove":
			fmt.Print("Input: ")
			input, err := readInput(reader)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			if removeItem(items, input) {
				fmt.Println("Removed!")
			} else {
				fmt.Println("Item not found.")
			}

		case "list":
			list := toList(items)
			if len(list) == 0 {
				fmt.Println("No items found.")
				continue
			}

			fmt.Println("Items:")
			for _, item := range list {
				fmt.Println("-", item)
			}

		case "quit":

			return

		default:
			fmt.Println("Invalid command.")
		}
	}
}

func readInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func addItem(items map[string]bool, input string) {
	items[input] = true
}

func checkItem(items map[string]bool, input string) bool {
	return items[input]
}

func removeItem(items map[string]bool, input string) bool {
	if items[input] {
		delete(items, input)
		return true
	}
	return false
}

func toList(items map[string]bool) []string {
	return slices.Collect(maps.Keys(items))
}
