package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func meminfo() ([]byte, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return io.ReadAll(file)
}

func printMeminfo() {
	b, err := meminfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}

func memdiff(fn func()) {
	r, _ := regexp.Compile(`^.*?:\s*(\d+)\s*(kB)?$`)

	b1, err := meminfo()
	if err != nil {
		log.Fatal(err)
	}
	fn()
	b2, err := meminfo()
	if err != nil {
		log.Fatal(err)
	}
	s1 := bufio.NewScanner(bytes.NewReader(b1))
	s2 := bufio.NewScanner(bytes.NewReader(b2))
	for s1.Scan() && s2.Scan() {
		a := s1.Text()
		u1 := r.ReplaceAllString(a, "$2")
		v1 := r.ReplaceAllString(a, "$1")
		vv1, err := strconv.Atoi(v1)
		if err != nil {
			log.Fatal(err)
		}
		b := s2.Text()
		u2 := r.ReplaceAllString(b, "$2")
		v2 := r.ReplaceAllString(b, "$1")
		vv2, err := strconv.Atoi(v2)
		if err != nil {
			log.Fatal(err)
		}
		if u1 != u2 {
			log.Fatal(fmt.Errorf("unit is not equal: %q != %q", u1, u2))
		}
		difference := vv2 - vv1
		var es1, es2 string
		if a != b {
			es1 = "\x1b[31m"
			es2 = "\x1b[0m"
		}
		fmt.Printf("%s%s\t->\t%s\t=\t%d %s%s\n", es1, a, b, difference, u1, es2)
	}
}
