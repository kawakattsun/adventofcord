package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const year = 2020

func main() {
	flag.Parse()
	part := flag.Arg(0)
	if part != "1" && part != "2" {
		log.Fatal("set part number. (1|2)")
	}

	fp, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var list []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		list = append(list, num)
	}

	if part == "1" {
		for _, i := range list {
			for _, j := range list {
				if i+j == year {
					fmt.Println(i * j)
					os.Exit(0)
				}
			}
		}
	} else if part == "2" {
		for _, i := range list {
			for _, j := range list {
				for _, k := range list {
					if i+j+k == year {
						fmt.Println(i * j * k)
						os.Exit(0)
					}
				}
			}
		}
	}
}
