// Copyright 2014 Karan Misra.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	port   = flag.Int("p", 5000, "port to serve on")
	prefix = flag.String("x", "/", "prefix to serve under")
)

func main() {
	flag.Parse()

	// Get the dir to serve
	if flag.NArg() < 1 {
		log.Fatalln("Please provide the dir to serve as the last argument. A simple . will also do")
	}
	dir := flag.Arg(0)
	portStr := fmt.Sprintf(":%v", *port)
	if !strings.HasSuffix(*prefix, "/") {
		*prefix = *prefix + "/"
	}

	log.Printf("Service traffic from %v under port %v with prefix %v\n", dir, *port, *prefix)
	log.Printf("Or simply put, just open http://localhost:%v%v to get rocking!\n", *port, *prefix)

	http.Handle(*prefix, http.StripPrefix(*prefix, http.FileServer(http.Dir(dir))))
	http.ListenAndServe(portStr, nil)
}
