package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
)

type RequestTracker struct {
	mu       sync.Mutex
	requests [][]byte
}

func (rt *RequestTracker) Track(req *http.Request) {
	rt.mu.Lock()
	defer rt.mu.Unlock()
	// alloc 10KB for each track
	rt.requests = append(rt.requests, bytes.Repeat([]byte("a"), 10000))
}

var (
	requests RequestTracker
)

func init() {
	requests = RequestTracker{}
}

func read(w http.ResponseWriter, r *http.Request) {
	requests.Track(r)
	r.Body.Close()
	w.WriteHeader(http.StatusOK)
	f, err := os.Open("/proc/1/status")
	if err != nil {
		fmt.Println(err)
	}
	stat := fmt.Sprint(f.Stat())
	w.Write([]byte(stat))
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func testBreak(w http.ResponseWriter, r *http.Request) {
	i := 0
	for i < 1000 {
		if fileExists("/etc/test_break") {
			fmt.Println("File exists, breaking.")
			break
		} else {
			fmt.Println("File not created yet.")
		}
		i = i + 1
	}
}

func write(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create("/etc/test_break")
    if err != nil {
        fmt.Println(err)
        return
    }
	l, err := f.WriteString("Hello World")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
	requests.Track(r)
	r.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprint(l, "bytes written successfully")))
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }

}

func main() {
	fmt.Printf("Starting Server on port :%s\n", "8080")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		requests.Track(r)
		r.Body.Close()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`OK`))
	})

	http.HandleFunc("/read", read)
	http.HandleFunc("/break", testBreak)
	http.HandleFunc("/write", write)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
