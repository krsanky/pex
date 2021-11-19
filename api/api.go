package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go.d34d.net/pex/fib"
	"go.d34d.net/pex/lg"
	"go.d34d.net/pex/session"
)

//  0-based index into the fibonacci sequence.
var FibIdxSessKey string = "pex-fib-idx"

func AddRoutes(m *http.ServeMux) {
	m.HandleFunc("/api", testIndex)
	m.HandleFunc("/current", current)
	m.HandleFunc("/next", next)
	m.HandleFunc("/previous", previous)
}

func testIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// test session
	skey := "pex-fib"
	sstr := session.Session.GetString(r.Context(), skey)
	lg.Log.Printf("sstr:%s", sstr)
	if sstr == "" {
		session.Session.Put(r.Context(), skey, "1")
	} else {
		i, err := strconv.Atoi(sstr)
		if err != nil {
			lg.Log.Printf("pex.Index ERR:%v", err)
		} else {
			i += i
			session.Session.Put(r.Context(), skey, fmt.Sprintf("%d", i))
		}
	}

	mp := make(map[string]interface{})
	mp["pex:testIndex"] = true
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}

// Get the current index and corresponding fibonacci
// sequencei value.
func current(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idx := session.Session.GetInt(r.Context(), FibIdxSessKey)
	// There is no practical distinction between no
	// session and session at 0.

	mp := make(map[string]interface{})
	mp["idx"] = idx
	mp["fib"] = fib.FibonacciLoop(idx)
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}

// Increment sequence index value and get the corresponding value.
func next(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idx := session.Session.GetInt(r.Context(), FibIdxSessKey)
	idx += 1
	fib := fib.FibonacciLoop(idx)
	lg.Log.Printf("fib:%d", fib)
	session.Session.Put(r.Context(), FibIdxSessKey, idx)

	mp := make(map[string]interface{})
	mp["idx"] = idx
	mp["fib"] = fib
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}

// Decrement sequence index value and get the corresponding value.
// If the index would get negative, then set values to -1
// to indicate error, and don't decrement.
// Index will not go below zero.
func previous(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idx := session.Session.GetInt(r.Context(), FibIdxSessKey)
	idx -= 1
	if idx < -1 {
		idx = -1
	}
	fib := fib.FibonacciLoop(idx)
	lg.Log.Printf("fib:%d", fib)
	if idx >=  0 {
		session.Session.Put(r.Context(), FibIdxSessKey, idx)
	}

	mp := make(map[string]interface{})
	mp["idx"] = idx
	mp["fib"] = fib
	bs, _ := json.Marshal(mp)

	w.Write(bs)
}
