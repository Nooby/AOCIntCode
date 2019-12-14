package vm

import (
	"fmt"
	"io"
	"log"
)

type VM struct {
	Mem   []int
	Pc    int
	modes []int
	run   bool
}

func (v *VM) Load(r io.Reader) error {
	var i int
	var mem []int
	c := true
	for c {
		_, err := fmt.Fscanf(r, "%d,", &i)
		if err != nil {
			c = false
			log.Printf("VM Error in Load at pos %v, %v", len(mem), err)
		}
		mem = append(mem, i)
	}
	v.Mem = mem

	fmt.Printf("VM Load %v ints\n", len(mem))
	return nil
}

func (v *VM) Patch(o int, r io.Reader) error {
	var i int
	for true {
		_, err := fmt.Fscanf(r, "%d,", &i)
		v.Mem[o] = i
		if err != nil {
			log.Printf("VM Error in Patch, %v", err)
			break
		}
		o++
	}
	return nil
}

func (v *VM) Run() {
	v.Pc = 0
	v.run = true
	for v.run {
		v.Step()
	}
}

func (v *VM) Step() bool {
	c := v.readOp()
	f := Ops[c]
	a := v.readArgs(f)

	f.Op(v, a)

	v.Pc += f.Args + 1
	return true
}

func (v *VM) readOp() int {
	op := v.Mem[v.Pc]
	if op == add || op == mult {
		v.modes = []int{0, 0, 1}
	}
	return op
}

func (v *VM) readArgs(code OpInfo) []int {
	args := v.Mem[v.Pc+1 : v.Pc+code.Args+1]
	var ra []int
	for i := 0; i < code.Args; i++ {
		if v.modes[i] == 0 {
			ra = append(ra, v.Mem[args[i]])
		} else if v.modes[i] == 1 {
			ra = append(ra, args[i])
		}
	}
	return ra
}
