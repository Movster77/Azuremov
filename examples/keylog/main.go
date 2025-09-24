package main

import (
	"fmt"
	"log"

	azuretls "github.com/Noooste/azuretls-client"
)

func main() {
	session := azuretls.NewSession()
	defer session.Close()

	session.Browser = azuretls.Chrome

	resp, err := session.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	negotiated := "unknown"
	if resp.HttpResponse != nil {
		if resp.HttpResponse.TLS != nil && resp.HttpResponse.TLS.NegotiatedProtocol != "" {
			negotiated = resp.HttpResponse.TLS.NegotiatedProtocol
		} else if resp.HttpResponse.Proto != "" {
			negotiated = resp.HttpResponse.Proto
		}
	}

	fmt.Printf("Negotiated protocol: %s\n", negotiated)
	fmt.Printf("Status code: %d\n", resp.StatusCode)
}
