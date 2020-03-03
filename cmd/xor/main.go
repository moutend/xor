package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	silentFlag := flag.Bool("silent", false, "don't output anything")
	replaceFlag := flag.Bool("replace", false, "don't leave original file(s)")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		return nil
	}
	for _, path := range args {
		out, err := xor(path)

		if err != nil {
			return err
		}
		if *replaceFlag {
			os.Remove(path)
		}
		if !*silentFlag {
			fmt.Printf("xor\t%s\t%s\n", filepath.Base(path), filepath.Base(out))
		}
	}

	return nil
}

func xor(src string) (dst string, err error) {
	if strings.HasSuffix(src, ".bin") {
		dst = strings.TrimSuffix(src, ".bin")
	} else {
		dst = src + ".bin"
	}

	data, err := ioutil.ReadFile(src)

	if err != nil {
		return dst, err
	}

	result := make([]byte, len(data))

	for i, _ := range data {
		result[i] = ^data[i]
	}
	if err := ioutil.WriteFile(dst, result, 0644); err != nil {
		return dst, err
	}

	return dst, nil
}
