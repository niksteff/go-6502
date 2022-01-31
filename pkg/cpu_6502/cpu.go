/*
 this is a software implementation of the 6502 microprocessor written in the go
 programming language
*/

package cpu6502

import "log"

// CPU represents a physical 6502 CPU
type CPU struct {
	PC Word // program counter
	AC Byte // accumulator
	X  Byte // x register
	Y  Byte // y register
	SP Byte // stack pointer

	SR Byte // status register, flags are used to set the values
}

// Reset will reset the CPU into a booting state
func (c *CPU) Reset(m *Mem) {
	log.Println("resetting cpu ...")

	m.init(MEM_CAP_apple2)

	// reset the stackpointer to 0xFF
	c.SP = c.SP - 0x01

	log.Println("cpu reset")
}

func (c *CPU) Execute(m *Mem, cycles uint32) {
	for cycles > 0 {
		var ins Byte = c.Fetch(m, &cycles)
		switch ins {
		case INS_LDA_IM:
			var val Byte = c.Fetch(m, &cycles)
			c.AC = val

			if c.AC == 0 {
				setBit(uint32(c.SR))
			}

			if c.AC & 0b10000000 > 0 {
				c.SR+=Byte(N)
			}
		}
	}
}

func (c *CPU) Fetch(m *Mem, cycles *uint32) Byte {
	var ins Byte = m.Access(c.PC)

	// inkrement the program counter to the next instruction
	c.PC++
	log.Printf("program counter: %d\n", c.PC)

	// we just consumed one cycle
	*cycles--

	return ins
}

// Here's a function to set a bit. First, shift the number 1 the specified
// number of spaces in the integer (so it becomes 0010, 0100, etc). Then OR it
// with the original input. This leaves the other bits unaffected but will
// always set the target bit to 1.
func setBit(n uint32, pos uint32) uint32 {
    n |= (1 << pos)
    return n
}

// Here's a function to clear a bit. First shift the number 1 the specified
// number of spaces in the integer (so it becomes 0010, 0100, etc). Then flip
// every bit in the mask with the ^ operator (so 0010 becomes 1101). Then use a
// bitwise AND, which doesn't touch the numbers AND'ed with 1, but which will
// unset the value in the mask which is set to 0.
func clearBit(n uint32, pos uint32) uint32 {
    var mask uint32 = ^(1 << pos)
    n &= mask
    return n
}

//Finally here's a function to check whether a bit is set. Shift the number 1
//the specified number of spaces (so it becomes 0010, 0100, etc) and then AND it
//with the target number. If the resulting number is greater than 0 (it'll be 1,
//2, 4, 8, etc) then the bit is set.
func hasBit(n uint32, pos uint32) bool {
    var val uint32 = n & (1 << pos)
    return (val > 0)
}