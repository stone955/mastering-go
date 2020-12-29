package options

import (
	"crypto/tls"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	s1, err := NewServer(":", 1024)
	if err != nil {
		t.Fatal(err)
	}
	s2, err := NewServer(":", 2048,
		Protocol("udp"),
	)
	if err != nil {
		t.Fatal(err)
	}
	s3, err := NewServer(":", 8080,
		DialTimeout(60*time.Second),
		MaxConn(1000),
		TLS(&tls.Config{}),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Server 1: %+v\n", *s1)
	t.Logf("Server 2: %+v\n", *s2)
	t.Logf("Server 3: %+v\n", *s3)

	/*
		=== RUN   TestNewServer
		    server_test.go:29: Server 1: {Addr:: Port:1024 Protocol:tcp DialTimeout:30s MaxConn:100 TLS:<nil>}
		    server_test.go:30: Server 2: {Addr:: Port:2048 Protocol:udp DialTimeout:30s MaxConn:100 TLS:<nil>}
		    server_test.go:31: Server 3: {Addr:: Port:8080 Protocol:tcp DialTimeout:1m0s MaxConn:1000 TLS:0xc000001b00}
		--- PASS: TestNewServer (0.00s)
		PASS
	*/
}
