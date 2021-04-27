package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	pr, pw := io.Pipe()
	defer func() {
		_ = pr.Close()
	}()

	f, err := os.Open("/Users/stone/Documents/Projects/go/stone/mastering-go/file/pipe/CDNPRDE04142A2012301002001.076.gz")
	if err != nil {
		log.Printf("os.Open() Error %v\n", err)
		return
	}
	defer func() {
		_ = f.Close()
	}()

	go func() {
		defer func() {
			_ = pw.Close()
		}()

		// 模拟等待建立连接
		<-time.After(5 * time.Second)

		wn, err := io.Copy(pw, f)
		if err != nil {
			log.Printf("io.Copy() Error %v\n", err)
			return
		}

		log.Printf("io.Copy() write %v bytes\n", wn)
	}()

	// 解压
	gr, err := gzip.NewReader(pr)
	if err != nil {
		log.Printf("gzip.NewReader() Error %v\n", err)
		return
	}

	var totals int

	for {
		p := make([]byte, 1024)
		n, err := gr.Read(p)
		if err != nil {
			if err == io.EOF {
				totals += n
				break
			}
			log.Printf("gzip.Read() Error %v\n", err)
			return
		}
		totals += n
	}

	log.Printf("Totals = %v\n", totals)

	/*
		// fatal error: all goroutines are asleep - deadlock!
		pr, _ := io.Pipe()
		defer func() {
			_ = pr.Close()
		}()

		// 解压
		_, err := gzip.NewReader(pr)
		if err != nil {
			log.Printf("gzip.NewReader() Error %v\n", err)
			return
		}
	*/
}
