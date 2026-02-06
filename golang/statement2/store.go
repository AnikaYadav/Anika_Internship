package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	items := make(map[string]bool)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Command (add, list, remove, check, quit):")
		fmt.Print("> ")

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(strings.ToLower(command))

		switch command {

		case "add":
			fmt.Print("Input: ")
			input := readInput(reader)
			addItem(items, input)
			fmt.Println("Added!")

		case "check":
			fmt.Print("Input: ")
			input := readInput(reader)

			if checkItem(items, input) {
				fmt.Println("Exists!")
			} else {
				fmt.Println("Does not exist.")
			}

		case "remove":
			fmt.Print("Input: ")
			input := readInput(reader)

			if removeItem(items, input) {
				fmt.Println("Removed!")
			} else {
				fmt.Println("Item not found.")
			}

		case "list":
			list := listItems(items)
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

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
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

func listItems(items map[string]bool) []string {
	result := []string{}
	for item := range items {
		result = append(result, item)
	}
	return result
}
