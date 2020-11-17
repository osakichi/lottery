package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	var list []string
	var max int

	flag.Parse()
	if flag.NArg() < 1 {
		log.Println("Usage:", os.Args[0], "<data_file>")
		os.Exit(1)
	}

	f, e := os.Open(flag.Arg(0))
	if e != nil {
		log.Print("os.Open(): ", e)
		os.Exit(2)
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	for r.Scan() {
		list = append(list, r.Text())
		if max < len(list[len(list)-1]) {
			max = len(list[len(list)-1])
		}
	}

	rand.Seed(time.Now().UnixNano())
	var choice int
	fmt.Println()
	for len(list) > 0 {
		for n := time.Duration(1000000); n <= time.Second; n *= 5 {
			for m := 0; m < int(time.Second/n); m++ {
				//fmt.Printf("%s\r", mix(list, max))
				choice = rand.Intn(len(list))
				fmt.Printf("\x1b\x5b\x41                    \r%s\n", list[choice])
				time.Sleep(n)
			}
			if len(list) == 1 {
				break
			}
		}

		fmt.Printf("\x1b\x5b\x41                    \r%s\n\n", list[choice])
		list = append(list[:choice], list[choice+1:]...)
	}
}

func mix(list []string, max int) string {
	var result string
	for n := 1; n < max; n++ {
		s := list[rand.Intn(len(list)-1)]
		if len(s) < n {
			result += " "
		} else {
			result += string(s[n-1])
		}
	}
	return result
}
