// Copyright 2014 Karan Misra.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//
// http authentication support added. References:
//   http://github.com/abbot/go-http-auth
//   http://stackoverflow.com/questions/25552107/golang-how-to-serve-static-files-with-basic-authentication

package main

import (
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/abbot/go-http-auth"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var version = "0.2.4"

var (
	port        = flag.Int("p", 5000, "port to serve on")
	prefix      = flag.String("x", "/", "prefix to serve under")
	showVersion = flag.Bool("v", false, "show version info")
	openBrowser = flag.Bool("o", false, "open the url")
	httpAuth    = flag.Bool("a", false, "requires random http auth")
	password    = "secretpassword"
	username    = "serve"
)

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func Secret(user, realm string) string {
	if user == username {
		return password
	}
	return ""
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println("serve version", version)
		os.Exit(0)
	}

	var dir string
	// Get the dir to serve
	if flag.NArg() < 1 {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Please provide the dir to serve as the last argument. A simple . will also do")
			os.Exit(1)
		}
		dir = cwd
	}
	dir = flag.Arg(0)
	portStr := fmt.Sprintf(":%v", *port)
	if !strings.HasPrefix(*prefix, "/") {
		*prefix = "/" + *prefix
	}
	if !strings.HasSuffix(*prefix, "/") {
		*prefix = *prefix + "/"
	}

	uri := fmt.Sprintf("http://localhost:%v%v", *port, *prefix)

	fmt.Printf("Service traffic from %v under port %v with prefix %v\n", dir, *port, *prefix)
	fmt.Printf("Or simply put, just open %v to get rocking!\n", uri)
	if *httpAuth {
		username = RandomString(5)
		password = RandomString(5)
		h := sha1.New()
		h.Write([]byte(password))
		fmt.Printf("user: %v password: %v\n", username, password)
		password = "{SHA}" + base64.StdEncoding.EncodeToString(h.Sum(nil))
		// fmt.Println(password)
	}
	go func() {
		if *openBrowser {
			success := waitForWebserver()
			if !success {
				// We have waited too long for the webserver to start; bail.
				fmt.Fprintf(os.Stderr, "The webserver did not start within the required time. Cannot open the browser for you\n")
				return
			}
			fmt.Printf("Opening your browser to %v\n", uri)
			cmd := exec.Command("open", uri)
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Could not open the url in your browser\n%v\n", err)
			}
		}
	}()

	authenticator := auth.NewBasicAuthenticator("Serve Credentials", Secret)
	if *httpAuth {
		http.HandleFunc("/", auth.JustCheck(authenticator, handleFileServer(dir, "/")))
	} else {
		http.HandleFunc("/", handleFileServer(dir, "/"))
	}
	if err := http.ListenAndServe(portStr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error while starting the web server\n%v\n", err)
		os.Exit(1)
	}
}

func handleFileServer(dir, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.URL)
		realHandler(w, req)
	}
}

func waitForWebserver() bool {
	timeout := time.After(1 * time.Second)
	connStr := fmt.Sprintf("127.0.0.1:%v", *port)
	for {
		select {
		case <-timeout:
			return false
		default:
			conn, err := net.DialTimeout("tcp", connStr, 50*time.Millisecond)
			if err != nil {
				continue
			}
			conn.Close()
			return true
		}
	}
}
