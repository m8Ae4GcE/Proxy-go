package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	listenPort := flag.Int("l", 80, "listen on port")
	redirectPort := flag.Int("r", 443, "redirect on port")

	flag.Parse()
	fmt.Println("listen on port:", *listenPort)
	fmt.Println("redirect on port:", *redirectPort)

	u := fmt.Sprint("http://localhost:",*redirectPort)

	urlParse, err := url.Parse(u)
	if err != nil {
		panic(err)
	}

	director := func(req *http.Request) {
		req.URL.Scheme = urlParse.Scheme
		req.URL.Host = urlParse.Host
	}

	reverseProxy := &httputil.ReverseProxy{Director: director}
	handler := handler{proxy: reverseProxy}

	http.Handle("/", handler)

	if *listenPort == 443 {
		http.ListenAndServeTLS(fmt.Sprintf(":%d", *listenPort), "localhost.pem", "localhost-key.pem", handler)
	} else {
		http.ListenAndServe(fmt.Sprintf(":%d", *listenPort), nil)
	}
}

type handler struct {
	proxy *httputil.ReverseProxy
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	h.proxy.ServeHTTP(w, r)
}