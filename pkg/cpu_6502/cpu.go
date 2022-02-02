/*
 this is a software implementation of the 6502 microprocessor written in the go
 programming language
*/

package cpu6502

import (
	"log"
)

// every flag has a bit inside the status register this ordered from bit 7 to
// bit 0. We do not have single bits in golang so to save space we will create
// these constants which will each contain the bit to set for the flag. If the
// bit will not be set we can just not add the value to the status register
const (
	N  Byte = 0b1000_0000 // Negative flag
	V  Byte = 0b0100_0000 // Overflow flag
	_I Byte = 0b0010_0000 // ignored flag
	B  Byte = 0b0001_0000 // Break flag
	D  Byte = 0b0000_1000 // Decimal flag (use BCD for arithmetics)
	I  Byte = 0b0000_0100 // Interrupt flag (IRQ disable)
	Z  Byte = 0b0000_0010 // Zero flag
	C  Byte = 0b0000_0001 // Carry flag

	MEM_CAP_apple2 uint32 = 1024 * 64 // set the max memory to 64KiB
)

// Byte represents an unsigned 8 bit integer
// TODO: decide whether to use gos byte alias
type Byte uint8

// Word represents two Bytes
type Word uint16

// Cycles is a cycle
type Cycles uint32

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

	// set the program counter to the start location
	c.PC = 0xFFFC

	// reset the stackpointer to 0xFF
	c.SP = 0xFF

	log.Println("cpu reset")
}

func (c *CPU) Execute(mem *Mem, cycles Cycles) {
	for cycles > 0 {
		var ins Byte = c.fetchByte(&cycles, mem, c.PC)
		switch ins {
		case INS_LDA_IM:
			// c2
			c.AC = c.fetchByte(&cycles, mem, c.PC)
			c.set_sr_flags_lda()
		case INS_LDA_ZP:
			// c2
			var zeroPageAddr Byte = c.fetchByte(&cycles, mem, c.PC)
			// c3
			c.AC = c.read(&cycles, mem, zeroPageAddr)
			c.set_sr_flags_lda()
		case INS_LDA_ZPX:
			// c2
			var zeroPageAddr Byte = c.fetchByte(&cycles, mem, c.PC)
			// c3
			addr := cycles.add(mem, zeroPageAddr, c.X)
			// c4
			c.AC = c.read(&cycles, mem, addr)
		case INS_JMP_ABS:
			// c2
			var adl Byte = c.fetchByte(&cycles, mem, c.PC)
			// c3
			var adh Byte = c.fetchByte(&cycles, mem, c.PC)
			c.PC = ((Word(adh) << 8) | Word(adl))
		case INS_JSR_ABS:
			// c2 fetch adl
			var adl Byte = c.fetchByte(&cycles, mem, c.PC)
			// c3 internal op, dummy cycle disregard the read immediately
			var _ Byte = mem.Read(Word(c.SP))
			cycles--
			// c4 push pch to stack
			// pch -> S
			mem.Write(0x0100 | Word(c.SP), Byte(c.PC >> 8))
			c.SP--
			cycles--
			// c5 push pcl to stack
			// pcl -> S
			mem.Write(0x0100 | Word(c.SP), Byte(c.PC))
			c.SP--
			cycles--
			// c6
			var adh Word = Word(mem.Read(c.PC)) << 8
			c.PC = adh | Word(adl)
			cycles--
		default:
			log.Printf("instruction not handled: %d\n", ins)
		}
	}
}

func (c *CPU) set_sr_flags_lda() {
	if c.AC == 0 {
		// set the zero flag
		c.SR.SetFlag(Z)
	}

	if c.AC.HasFlag(N) {
		// set the negative flag
		c.SR.SetFlag(N)
	}
}

// fetchByte returns the next instruction from the memory
// it increases the program counter by 1
// it decreases the cycles by 1
func (c *CPU) fetchByte(cycles *Cycles, m *Mem, addr Word) Byte {
	var b Byte = m.Read(addr)
	c.PC++
	*cycles--
	return b
}

func (cycles *Cycles) Push(m *Mem, c *CPU, b Byte) {
	// the address of the sp is always between 0x01FF and 0x0100
	// the stack decrements when growing
	// after reaching 0x100 the sp overflows back to 0x01FF
	m.Write(0x0100 | Word(c.SP), b)
	c.SP--
	*cycles--
}

// add will add up two addresses and returns the result
// this can overflow and the resulting address will warp around
func (cycles *Cycles) add(m *Mem, a Byte, b Byte) Byte {
	var addr Byte = a + b

	// we just consumed one cycle
	*cycles--

	return addr
}

// read will read a byte from memory but does not advance the program counter
// because we are not loading the next instruction
func (c *CPU) read(cycles *Cycles, m *Mem, addr Byte) Byte {
	var b Byte = m.Read(Word(addr))

	// we just consumed one cycle
	*cycles--

	return b
}

// SetFlag will set the given bit on the status register
func (sr *Byte) SetFlag(flag Byte) {
	*sr = (*sr | flag)
}

// UnsetFlag will unset the given bit on the status register
func (sr *Byte) UnsetFlag(flag Byte) {
	*sr &= (^flag)
}

// HasFlag checks if the given bit is set on the status register
func (sr *Byte) HasFlag(flag Byte) bool {
	return (*sr & flag) > 0
}
