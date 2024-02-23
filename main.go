package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const DEFAULT_PORT = 2880

func main() {
	port := flag.Int("p", DEFAULT_PORT, "Set The Http Port")
	flag.Parse()
	pwd, _ := os.Getwd()
	fmt.Println("srart httpserver......dir: "+pwd+", on ", *port)
	http.Handle("/", http.FileServer(http.Dir(pwd)))
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
