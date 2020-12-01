package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const year = 2020

func main() {
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
	for _, num := range list {
		for _, n := range list {
			if num+n == year {
				fmt.Println(num * n)
				os.Exit(0)
			}
		}
	}
}
