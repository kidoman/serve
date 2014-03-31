// Copyright 2014 Karan Misra.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var version = "0.1.4"

var (
	port        = flag.Int("p", 5000, "port to serve on")
	prefix      = flag.String("x", "/", "prefix to serve under")
	showVersion = flag.Bool("v", false, "show version info")
)

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println("serve version", version)
		os.Exit(0)
	}

	// Get the dir to serve
	if flag.NArg() < 1 {
		fmt.Println("Please provide the dir to serve as the last argument. A simple . will also do")
		os.Exit(1)
	}
	dir := flag.Arg(0)
	portStr := fmt.Sprintf(":%v", *port)
	if !strings.HasSuffix(*prefix, "/") {
		*prefix = *prefix + "/"
	}

	fmt.Printf("Service traffic from %v under port %v with prefix %v\n", dir, *port, *prefix)
	fmt.Printf("Or simply put, just open http://localhost:%v%v to get rocking!\n", *port, *prefix)

	http.Handle(*prefix, http.StripPrefix(*prefix, http.FileServer(http.Dir(dir))))
	if err := http.ListenAndServe(portStr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error while starting the web server\n%v\n", err)
		os.Exit(1)
	}
}
