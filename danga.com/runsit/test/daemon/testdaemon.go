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
	"net/http"
	"os"
	"strconv"
)

var port = flag.Int("port", 8000, "port")

func main() {
	fmt.Fprintf(os.Stderr, "Hello on stderr\n")
	fmt.Fprintf(os.Stdout, "Hello on stdout\n")
	http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stderr, "crashing")
		os.Exit(2)
	})
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}