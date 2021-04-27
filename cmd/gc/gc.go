package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func printMemStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	log.Printf("Mem Alloc: %v\n", mem.Alloc)
	log.Printf("Mem TotalAlloc: %v\n", mem.TotalAlloc)
	log.Printf("Mem HeapAlloc: %v\n", mem.HeapAlloc)
	log.Printf("Mem NumGC: %v\n", mem.NumGC)
	log.Println("====================")
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		log.Fatalln(err)
	}
	defer trace.Stop()

	var mem runtime.MemStats
	printMemStats(mem)

	for i := 0; i < 10; i++ {
		b := make([]byte, 50000000)
		if b == nil {
			log.Println("make []byte error")
		}
	}
	printMemStats(mem)

	for i := 0; i < 10; i++ {
		b := make([]byte, 100000000)
		if len(b) == 0 {
			log.Println("make []byte error")
		}
		time.Sleep(time.Millisecond)
	}
	printMemStats(mem)

	/*
		stone@stonedeMacBook-Pro gc % ./gc
		2020/12/16 22:31:14 Mem Alloc: 124144
		2020/12/16 22:31:14 Mem TotalAlloc: 124144
		2020/12/16 22:31:14 Mem HeapAlloc: 124144
		2020/12/16 22:31:14 Mem NumGC: 0
		2020/12/16 22:31:14 ====================
		2020/12/16 22:31:14 Mem Alloc: 117112
		2020/12/16 22:31:14 Mem TotalAlloc: 500174024
		2020/12/16 22:31:14 Mem HeapAlloc: 117112
		2020/12/16 22:31:14 Mem NumGC: 10
		2020/12/16 22:31:14 ====================
		2020/12/16 22:31:14 Mem Alloc: 119280
		2020/12/16 22:31:14 Mem TotalAlloc: 1500257040
		2020/12/16 22:31:14 Mem HeapAlloc: 119280
		2020/12/16 22:31:14 Mem NumGC: 20
		2020/12/16 22:31:14 ====================
	*/
}
