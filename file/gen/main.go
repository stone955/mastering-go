package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fsp := "/Users/stone/Documents/Projects/go/unicom/testdata/pfau/test_file/31/normal_md5/CMBDSDWAL2020111018017.031"
	msp := "/Users/stone/Documents/Projects/go/unicom/testdata/pfau/test_file/31/normal_md5/CMBDSDWAL2020111018017.031.MD5"
	for i := 0; i < 10000; i++ {
		if i < 10 {
			fdp := fmt.Sprintf("/Users/stone/Documents/Projects/go/unicom/testdata/pfau/test_file/31/more_normal_md5/CMBDSDWAL202111101801%v.031", i)
			if err := CopyFile(fdp, fsp); err != nil {
				panic(err)
			}

			mdp := fmt.Sprintf("/Users/stone/Documents/Projects/go/unicom/testdata/pfau/test_file/31/more_normal_md5/CMBDSDWAL202111101801%v.031.MD5", i)
			if err := CopyFile(mdp, msp); err != nil {
				panic(err)
			}
			continue
		}
		if i < 100 {
			fdp := fmt.Sprintf("/Users/stone/Documents/Projects/go/unicom/testdata/pfau/test_file/31/more_normal_md5/CMBDSDWAL20211110180%v.031", i)
			if err := CopyFile(fdp, fsp); err != nil {
				panic(err)
			}

			mdp := fmt.Sprintf("/Users/stone/Documents/Projects/go/unicom/testdata/pfau/test_file/31/more_normal_md5/CMBDSDWAL20211110180%v.031.MD5", i)
			if err := CopyFile(mdp, msp); err != nil {
				panic(err)
			}
			continue
		}
		if i < 1000 {
			fdp := fmt.Sprintf("/Users/stone/Documents/Projects/go/unicom/testdata/pfau/test_file/31/more_normal_md5/CMBDSDWAL2021111018%v.031", i)
			if err := CopyFile(fdp, fsp); err != nil {
				panic(err)
			}

			mdp := fmt.Sprintf("/Users/stone/Documents/Projects/go/unicom/testdata/pfau/test_file/31/more_normal_md5/CMBDSDWAL2021111018%v.031.MD5", i)
			if err := CopyFile(mdp, msp); err != nil {
				panic(err)
			}
			continue
		}
		if i < 10000 {
			fdp := fmt.Sprintf("/Users/stone/Documents/Projects/go/unicom/testdata/pfau/test_file/31/more_normal_md5/CMBDSDWAL202111101%v.031", i)
			if err := CopyFile(fdp, fsp); err != nil {
				panic(err)
			}

			mdp := fmt.Sprintf("/Users/stone/Documents/Projects/go/unicom/testdata/pfau/test_file/31/more_normal_md5/CMBDSDWAL202111101%v.031.MD5", i)
			if err := CopyFile(mdp, msp); err != nil {
				panic(err)
			}
			continue
		}
	}
}

func CopyFile(dstPath, srcPath string) error {
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
