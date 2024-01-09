package main

import (
  "crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

const timeFmt = "2006-01-02T15:04:05.000Z"

func main() {
	req, _ := http.NewRequest("GET", "https://www.google.com", nil)
	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			fmt.Printf("%s: Get Conn: %s\n", time.Now().Format(timeFmt), hostPort)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("%s: Got Conn: %+v\n", time.Now().Format(timeFmt), connInfo)
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			fmt.Printf("%s: DNS Start: %+v\n", time.Now().Format(timeFmt), info)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("%s: DNS Done: %+v\n", time.Now().Format(timeFmt), dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Printf("%s: Connect Start: netework=%s, address=%s\n", time.Now().Format(timeFmt), network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Printf("%s: Connect Done: netework=%s, address=%s, err=%+v\n", time.Now().Format(timeFmt), network, addr, err)
		},
		TLSHandshakeStart: func() {
			fmt.Printf("%s: TLS Handshake Start\n", time.Now().Format(timeFmt))
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			fmt.Printf("%s: TLS Handshake Done: state=%+v, err=%+v\n", time.Now().Format(timeFmt), state, err)
		},
    WroteHeaders: func() {
      fmt.Printf("%s: Wrote Headers\n", time.Now().Format(timeFmt))
    },
    WroteRequest: func(info httptrace.WroteRequestInfo) {
      fmt.Printf("%s: Wrote Request: %+v\n", time.Now().Format(timeFmt), info)
    },
		GotFirstResponseByte: func() {
			fmt.Printf("%s: Got 1st Response Byte\n", time.Now().Format(timeFmt))
		},
    Wait100Continue: func() {
      fmt.Printf("%s: Wait 100 Continue\n", time.Now().Format(timeFmt))
    },
    Got100Continue: func() {
      fmt.Printf("%s: Got 100 Continue\n", time.Now().Format(timeFmt))
    },
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
  req.Header.Set("Expect", "100-continue")
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
