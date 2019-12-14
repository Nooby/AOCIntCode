package vm

type Op func(v *VM, args []int)

type OpInfo struct {
	Op   Op
	Args int
}

const (
	add  = 1
	mult = 2
	halt = 99
)

var Ops = map[int]OpInfo{
	add:  OpInfo{addFunc, 3},
	mult: OpInfo{multFunc, 3},
	halt: OpInfo{haltFunc, 0},
}

// add a b res
// adds a to b and stores the result in res
func addFunc(v *VM, args []int) {
	v.Mem[args[2]] = args[0] + args[1]
}

// mult a b res
// multiplies a with b and stores the result in res
func multFunc(v *VM, args []int) {
	v.Mem[args[2]] = args[0] * args[1]
}

// halt
// halts program execution
func haltFunc(v *VM, args []int) {
	v.run = false
}
