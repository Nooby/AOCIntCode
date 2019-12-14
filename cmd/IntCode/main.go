package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	vm "github.com/Nooby/AOCIntCode"
)

func main() {
	p1 := vm.VM{}

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
		case "patch":
			i, _ := strconv.Atoi(comm[1])
			r := strings.NewReader(comm[2])
			p1.Patch(i, r)
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
