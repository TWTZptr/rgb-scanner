package main

import (
	"flag"
	"fmt"
	"net/url"
	rbts "rgb-scanner/internal/robotstxt"
)

func main() {
	flag.Parse()
	unnamedArgs := flag.Args()

	if len(unnamedArgs) == 0 {
		fmt.Println("No target url")
		return
	}

	targetUrl := unnamedArgs[0]

	u, err := url.ParseRequestURI(targetUrl)

	if err != nil {
		fmt.Printf("URI was not parsed %s\n", targetUrl)
		return
	}

	rbts.ScanRobotsTxt(u.Host)
}
