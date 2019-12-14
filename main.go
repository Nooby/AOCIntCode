package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1 := VM{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		comm := strings.Split(scanner.Text(), " ")
		switch comm[0] {
		case "load":
			r, _ := os.Open("input/" + comm[1])
			p1.Load(r)
		case "read":
			r := strings.NewReader(comm[1])
			p1.Load(r)
		case "mem":
			if len(comm) == 1 {
				fmt.Printf("%v\n", p1.Mem)
			} else if len(comm) == 2 {
				i, _ := strconv.Atoi(comm[1])
				fmt.Printf("%v\n", p1.Mem[i])
			}
		case "run":
			p1.Run()
		}
		fmt.Print("> ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
