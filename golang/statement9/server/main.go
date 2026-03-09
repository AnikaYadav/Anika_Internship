package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"

	pb "trie-grpc/proto"
	"trie-grpc/trie"

	"google.golang.org/grpc"
)

const dataPath = "/data/trie_db.txt"

func loadTrie(t *trie.Trie) {
	file, err := os.Open(dataPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("No existing data file found, starting fresh.")
			return
		}
		log.Fatalf("failed to open storage: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if word != "" {
			t.Insert(word)
		}
	}
	log.Println("Trie loaded from storage.")
}

type server struct {
	pb.UnimplementedTrieServiceServer
	t *trie.Trie
}

func (s *server) Add(ctx context.Context, w *pb.Word) (*pb.Status, error) {
	s.t.Insert(w.Text)

	f, err := os.OpenFile(dataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return &pb.Status{Ok: false}, err
	}
	defer f.Close()

	if _, err := f.WriteString(w.Text + "\n"); err != nil {
		return &pb.Status{Ok: false}, err
	}

	return &pb.Status{Ok: true}, nil
}

func (s *server) Check(ctx context.Context, w *pb.Word) (*pb.Status, error) {
	ok := s.t.Search(w.Text)
	return &pb.Status{Ok: ok}, nil
}

func (s *server) Remove(ctx context.Context, w *pb.Word) (*pb.Status, error) {
	ok := s.t.Delete(w.Text)
	if !ok {
		return &pb.Status{Ok: false}, nil
	}

	f, err := os.OpenFile(dataPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return &pb.Status{Ok: false}, err
	}
	defer f.Close()

	for _, word := range s.t.ListAll() {
		f.WriteString(word + "\n")
	}

	return &pb.Status{Ok: true}, nil
}

func (s *server) List(ctx context.Context, e *pb.Empty) (*pb.WordList, error) {
	return &pb.WordList{Words: s.t.ListAll()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	t := trie.NewTrie()
	loadTrie(t)

	grpcServer := grpc.NewServer()
	pb.RegisterTrieServiceServer(grpcServer, &server{t: t})

	log.Println("Server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
