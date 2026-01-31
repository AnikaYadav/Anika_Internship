package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func initRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	fmt.Println("Connected to Redis:", redisAddr)
}

func getBaseURL() string {
	base := os.Getenv("BASE_URL")
	if base == "" {
		base = "http://localhost:8080"
	}
	return base
}

func generateShortID() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := make([]byte, 6)
	for i := range id {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	id := generateShortID()

	err = rdb.Set(ctx, id, req.URL, 0).Err()
	if err != nil {
		http.Error(w, "Redis write failed", http.StatusInternalServerError)
		return
	}

	shortURL := fmt.Sprintf("%s/%s", getBaseURL(), id)

	resp := ShortenResponse{ShortURL: shortURL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]

	val, err := rdb.Get(ctx, id).Result()
	if err == redis.Nil {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Redis read error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, val, http.StatusFound)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	initRedis()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
