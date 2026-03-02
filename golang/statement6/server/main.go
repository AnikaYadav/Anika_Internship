package main

import (
	"context"
	"log"
	"net"

	pb "trie-grpc/proto"
	"trie-grpc/trie"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTrieServiceServer
	t *trie.Trie
}

func (s *server) Add(ctx context.Context, w *pb.Word) (*pb.Status, error) {
	s.t.Insert(w.Text)
	return &pb.Status{Ok: true}, nil
}

func (s *server) Check(ctx context.Context, w *pb.Word) (*pb.Status, error) {
	ok := s.t.Search(w.Text)
	return &pb.Status{Ok: ok}, nil
}

func (s *server) Remove(ctx context.Context, w *pb.Word) (*pb.Status, error) {
	ok := s.t.Delete(w.Text)
	return &pb.Status{Ok: ok}, nil
}

func (s *server) List(ctx context.Context, e *pb.Empty) (*pb.WordList, error) {
	return &pb.WordList{Words: s.t.ListAll()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTrieServiceServer(grpcServer, &server{
		t: trie.NewTrie(),
	})

	log.Println("Server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
