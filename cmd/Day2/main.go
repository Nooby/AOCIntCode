package main

import (
	"fmt"
	"os"

	vm "github.com/Nooby/AOCIntCode"
)

func main() {
	Part1()
	Part2()
}

func Part1() {
	p1 := vm.VM{}
	r, _ := os.Open("input/day2")
	p1.Load(r)

	p1.Mem[1] = 12
	p1.Mem[2] = 2

	p1.Run()

	fmt.Printf("Solution Part 1: %v\n", p1.Mem[0])
}

func Part2() {
	p1 := vm.VM{}
	r, _ := os.Open("input/day2")
	p1.Load(r)

	snap := make([]int, 137)
	copy(snap, p1.Mem)
	fmt.Printf("Memory snap %v\n", snap)

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(p1.Mem, snap)
			p1.Mem[1] = noun
			p1.Mem[2] = verb

			p1.Run()

			if p1.Mem[0] == 19690720 {
				fmt.Printf("Noun: %v\nVerb: %v\nSolution Part 2: %v\n", noun, verb, 100*noun+verb)
				return
			}
		}
	}
}
