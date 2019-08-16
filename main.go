package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/liuzl/gocc"
	"github.com/liuzl/goutil"
)

var (
	input  = flag.String("i", "", "Input direction")
	output = flag.String("o", "", "Out direction")
	config = flag.String("c", "", "s2t / s2tw / s2twp / s2hk / t2s / t2tw / tw2s / tw2sp / t2hk / hk2s")
)

func mapToFile(occ *gocc.OpenCC, src, dst string) {
	srcF, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	dstF, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	br := bufio.NewReader(srcF)
	if err != nil {
		log.Fatal(err)
	}

	goutil.ForEachLine(br, func(line string) error {
		str, e := occ.Convert(line)
		if e != nil {
			return e
		}
		fmt.Fprint(dstF, str+"\n")
		return nil
	})
}

func mapToDir(occ *gocc.OpenCC, src, dst string) {
	if fis, err := ioutil.ReadDir(src); err == nil {
		for _, f := range fis {
			srcPath := filepath.Join(src, f.Name())
			dstPath := filepath.Join(dst, f.Name())
			if _, err := os.Stat(dstPath); os.IsNotExist(err) {
				if f.IsDir() {
					os.Mkdir(dstPath, os.ModePerm)
					mapToDir(occ, srcPath, dstPath)
				} else {
					mapToFile(occ, srcPath, dstPath)
				}
			}
		}
	} else {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	if *config == "" {
		*config = "s2t"
	}

	if *input == "" {
		log.Fatal("The input option is necessary.")
	}

	if *output == "" {
		log.Fatal("The output option is necessary.")
	}

	s2t, err := gocc.New(*config)
	if err != nil {
		log.Fatal(err)
	}

	mapToDir(s2t, *input, *output)

}
