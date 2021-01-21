package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const (
	D0x01 uint8 = 0x01
)

// LastLineSplitCount 返回指定分隔符的文件最后一行字段个数
func LastLineSplitCount(filename string, delimit uint8) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}

	defer f.Close()

	var text []byte

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		text = sc.Bytes()
	}

	return SplitCount(text, delimit), nil
}

// SplitCount 返回指定分隔符的字段个数
func SplitCount(bs []byte, delimit uint8) int {
	if len(bs) == 0 {
		return 0
	}

	count := 1
	for _, b := range bs {
		if b == delimit {
			count++
		}
	}
	return count
}

func main() {
	filename := `/Users/stone/Documents/Projects/go/stone/mastering-go/file/D0104220201230010001.AVL`
	count, err := LastLineSplitCount(filename, D0x01)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The columns of the last line: %v\n", count)

	filename = `/Users/stone/Documents/Projects/go/stone/mastering-go/file/CDNPRDE04142A2012301001001.076`
	md5Str, err := MD5Sum(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The checksum is %v\n", md5Str)
}

func MD5Sum(filename string) (string, error) {
	// 每块多大
	chunk := int64(16)

	file, err := os.Open(filename)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	info, _ := file.Stat()

	filesize := info.Size()

	blocks := int64(math.Ceil(float64(filesize) / float64(chunk)))

	hash := md5.New()

	for i := int64(0); i < blocks; i++ {
		size := int(math.Min(float64(chunk), float64(filesize-i*chunk)))
		buf := make([]byte, size)
		if _, err := file.Read(buf); err != nil {
			return "", err
		}
		if _, err := io.WriteString(hash, string(buf)); err != nil {
			return "[]byte", err
		}
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
