package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/krsanky/pex/cookie"
	"github.com/krsanky/pex/fib"
	"github.com/krsanky/pex/lg"
)

func AddRoutes(m *http.ServeMux) {
	m.HandleFunc("/current", current)
	m.HandleFunc("/next", next)
	m.HandleFunc("/previous", previous)
	m.HandleFunc("/random", random)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Get the current index and corresponding fibonacci
// sequencei value.
func current(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idx := cookie.GetFibIdx(r)
	// There is no practical distinction between no
	// session and session at 0.

	mp := make(map[string]interface{})
	mp["idx"] = idx
	mp["fib"] = fib.MemoFib(idx)
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}

// Increment sequence index value and get the corresponding value.
func next(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idx := cookie.GetFibIdx(r)
	lg.Log.Printf("get cookie:%d", idx)
	idx += 1
	fib := fib.MemoFib(idx)
	lg.Log.Printf("fib := fib.MemoFib(idx):%d", idx)
	cookie.SetFibIdx(w, idx)

	mp := make(map[string]interface{})
	mp["idx"] = idx
	mp["fib"] = fib
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}

// Decrement sequence index value and get the corresponding value.
// If the index would go negative, then set values to -1
// to indicate error, and don't decrement.
// Index will not go below zero.
func previous(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idx := cookie.GetFibIdx(r)
	idx -= 1
	if idx < -1 {
		idx = -1
	}
	fib := fib.MemoFib(idx)
	lg.Log.Printf("fib:%d", fib)
	if idx >= 0 {
		cookie.SetFibIdx(w, idx)
	}

	mp := make(map[string]interface{})
	mp["idx"] = idx
	mp["fib"] = fib
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}

// This is for testing, without having to worry about cookies.
func random(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idx := rand.Intn(1000)
	fib := fib.MemoFib(idx)

	mp := make(map[string]interface{})
	mp["idx"] = idx
	mp["fib"] = fib
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}
