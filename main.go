package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	var (
		err                                                             error
		targetUrl, certificateFilePath, keyFilePath, listeningInterface string
		listeningPort                                                   uint
	)

	flag.StringVar(&targetUrl, "url", "", "[REQ] the target url to forward requests to")
	flag.StringVar(&certificateFilePath, "crt", "", "[REQ] certificate file")
	flag.StringVar(&keyFilePath, "key", "", "[REQ] key file")
	flag.StringVar(&listeningInterface, "if", "127.0.0.1", "[OPT] the interface to listen on")
	flag.UintVar(&listeningPort, "port", 443, "[OPT] the port to listen on")
	flag.Parse()

	if len(targetUrl) == 0 || len(certificateFilePath) == 0 || len(keyFilePath) == 0 {
		fmt.Fprintln(os.Stderr, "Not all required arguments were given. Use -h!")
		os.Exit(1)
	}

	parsedUrl, err := url.ParseRequestURI(targetUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	listenUrl := fmt.Sprintf("%s:%d", listeningInterface, listeningPort)

	http.Handle("/", httputil.NewSingleHostReverseProxy(parsedUrl))
	err = http.ListenAndServeTLS(listenUrl, certificateFilePath, keyFilePath, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
