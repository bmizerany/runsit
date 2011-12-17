/*
Copyright 2011 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var port = flag.Int("port", 8000, "port")

func crashHandler(w http.ResponseWriter, r *http.Request) {
	status := 2
	if st := r.FormValue("status"); st != "" {
		status, _ = strconv.Atoi(st)
	}
	fmt.Fprintf(os.Stderr, "crashing with status %d", status)
	os.Exit(status)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "pid=%d\n", os.Getpid())
	cwd, _ := os.Getwd()
	fmt.Fprintf(w, "cwd=%s\n", cwd)
	fmt.Fprintf(w, "uid=%d\n", os.Getuid())
	fmt.Fprintf(w, "euid=%d\n", os.Geteuid())
	fmt.Fprintf(w, "gid=%d\n", os.Getgid())

	for _, env := range os.Environ() {
		fmt.Fprintf(w, "%s\n", env)
	}

}

func logNoise() {
	for {
		log.Printf("some log noise")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Fprintf(os.Stdout, "Hello on stdout; listening on port %d\n", *port)
	fmt.Fprintf(os.Stderr, "Hello on stderr\n")
	go logNoise()
	http.HandleFunc("/crash", crashHandler)
	http.HandleFunc("/", statusHandler)
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
