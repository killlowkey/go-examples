package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // start pprof
	"sync"
)

type UserData struct {
	Data []byte
}

type UserCache struct {
	mu    sync.Mutex
	Cache map[string]*UserData
}

func NewUserCache() *UserCache {
	return &UserCache{
		Cache: make(map[string]*UserData),
	}
}

var userCache = NewUserCache()

func handleRequest(w http.ResponseWriter, r *http.Request) {
	userCache.mu.Lock()
	defer userCache.mu.Unlock()

	userData := &UserData{
		Data: make([]byte, 1000000),
	}

	userID := fmt.Sprintf("%d", len(userCache.Cache))
	userCache.Cache[userID] = userData
	log.Printf("Added data for user %s. Total users: %d\n", userID, len(userCache.Cache))
}

func main() {
	http.HandleFunc("/leaky-endpoint", handleRequest)
	http.ListenAndServe(":8080", nil)
}
