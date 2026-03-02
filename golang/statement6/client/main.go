package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	pb "trie-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func readInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func main() {

	conn, err := grpc.Dial("trie-server:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewTrieServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Commands: add, remove, check, list, quit")

	for {
		fmt.Print("\n> ")

		cmd, err := readInput(reader)
		if err != nil {
			fmt.Println("Error reading command:", err)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		switch strings.ToLower(cmd) {

		case "add":
			fmt.Print("Input: ")
			word, err := readInput(reader)
			if err != nil {
				fmt.Println("Error reading input:", err)
				cancel()
				continue
			}
			if word == "" {
				fmt.Println("Word cannot be empty.")
				cancel()
				continue
			}

			_, err = client.Add(ctx, &pb.Word{Text: word})
			if err != nil {
				fmt.Println("RPC error:", err)
				cancel()
				continue
			}

			fmt.Println("Added!")

		case "remove":
			fmt.Print("Input: ")
			word, err := readInput(reader)
			if err != nil {
				fmt.Println("Error reading input:", err)
				cancel()
				continue
			}

			res, err := client.Remove(ctx, &pb.Word{Text: word})
			if err != nil {
				fmt.Println("RPC error:", err)
				cancel()
				continue
			}

			if res.Ok {
				fmt.Println("Removed!")
			} else {
				fmt.Println("Item not found.")
			}

		case "check":
			fmt.Print("Input: ")
			word, err := readInput(reader)
			if err != nil {
				fmt.Println("Error reading input:", err)
				cancel()
				continue
			}

			res, err := client.Check(ctx, &pb.Word{Text: word})
			if err != nil {
				fmt.Println("RPC error:", err)
				cancel()
				continue
			}

			if res.Ok {
				fmt.Println("Exists!")
			} else {
				fmt.Println("Does not exist.")
			}

		case "list":
			res, err := client.List(ctx, &pb.Empty{})
			if err != nil {
				fmt.Println("RPC error:", err)
				cancel()
				continue
			}

			if len(res.Words) == 0 {
				fmt.Println("No items found.")
				cancel()
				continue
			}

			fmt.Println("Items:")
			for _, item := range res.Words {
				fmt.Println("-", item)
			}

		case "quit":
			cancel()
			return

		default:
			fmt.Println("Invalid command.")
		}

		cancel()
	}
}
