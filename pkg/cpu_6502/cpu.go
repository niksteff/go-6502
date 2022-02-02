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
	N  Byte = 0x80 // Negative flag
	V  Byte = 0x40 // Overflow flag
	_I Byte = 0x20 // ignored flag
	B  Byte = 0x10 // Break flag
	D  Byte = 0x8  // Decimal flag (use BCD for arithmetics)
	I  Byte = 0x4  // Interrupt flag (IRQ disable)
	Z  Byte = 0x2  // Zero flag
	C  Byte = 0x1  // Carry flag

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
			c.AC = c.read(&cycles, mem, Word(zeroPageAddr))
			c.set_sr_flags_lda()
		case INS_LDA_ZPX:
			// c2
			var zeroPageAddr Byte = c.fetchByte(&cycles, mem, c.PC)
			// c3
			var addr Byte = cycles.add(mem, zeroPageAddr, c.X)
			// c4
			c.AC = c.read(&cycles, mem, Word(addr))
		case INS_JMP_ABS:
			// c2
			var adl Byte = c.fetchByte(&cycles, mem, c.PC)
			// c3
			var adh Byte = c.fetchByte(&cycles, mem, c.PC)
			c.PC = ((Word(adh)<<8) | Word(adl))
		case INS_JSR:
			// c2 fetch adl
			var adl Byte = c.fetchByte(&cycles, mem, c.PC)
			// c3 internal dummy op
			var _ Byte = mem.Read(Word(c.SP))
			cycles--
			// c4 push pch to stack
			// pch -> S
			var t Word = c.PC >> 8
			var pch Byte = Byte(t)
			c.PushToStack(&cycles, mem, pch)
			// c5 push pcl to stack
			// pcl -> S
			var pcl Byte = Byte(c.PC)
			c.PushToStack(&cycles, mem, pcl)
			// c6
			c.PC = Word(mem.Read(c.PC))<<8 | Word(adl)
			cycles--
		case INS_RTS:
			// c2 internal dummy op
			var _ Byte = c.read(&cycles, mem, c.PC)
			// c3
			cycles--
			// c4 (s) -> pcl
			// var pcl Byte = mem.Read(0x0100|Word(c.SP))
			var pcl Byte = c.PopFromStack(&cycles, mem)
			// c5 (s) -> pch
			// var pch Byte = mem.Read(0x0100|Word(c.SP))
			var pch Byte = c.PopFromStack(&cycles, mem)
			c.PC = Word(pch)<<8 | Word(pcl)
			// c6
			c.PC++
			cycles--
		default:
			log.Printf("instruction not handled: %d\n", ins)
		}
	}
}

func (c *CPU) set_sr_flags_lda() {
	if c.AC == 0 {
		// set the zero flag
		c.SR.SetBit(Z)
	}

	if c.AC.HasBit(N) {
		// set the negative flag
		c.SR.SetBit(N)
	}
}

// fetchByte returns the next instruction from the memory
// pc advance = 1
// cycle cost = 1
func (c *CPU) fetchByte(cycles *Cycles, m *Mem, addr Word) Byte {
	var b Byte = m.Read(addr)
	c.PC++
	*cycles--
	return b
}

// PushToStrack writes the given byte to the next free stack address and sets
// the stack to the next free address in the page
// cycle cost = 1
func (c *CPU) PushToStack(cycles *Cycles, mem *Mem, b Byte) {
	// the address of the sp is always between 0x01FF and 0x0100
	// the stack decrements when growing
	// after reaching 0x100 the sp overflows back to 0x01FF
	mem.Write(0x0100|Word(c.SP), b)
	c.SP--
	*cycles--
}

// PopFromStack reads the last written Byte of the stack from memory and frees
// one byte from the stack page
// cycle cost = 1
func (c *CPU) PopFromStack(cycles *Cycles, mem *Mem) Byte {
	c.SP++

	var b Byte = mem.Read(0x0100 | Word(c.SP))

	*cycles--

	return b
}

// add will add up two addresses and returns the result
// this can overflow and the resulting address will warp around
// cycle cost = 1
func (cycles *Cycles) add(mem *Mem, a Byte, b Byte) Byte {
	var addr Byte = a + b

	*cycles--

	return addr
}

// read will read a byte from memory but does not advance the program counter
// cycle cost = 1
func (c *CPU) read(cycles *Cycles, mem *Mem, addr Word) Byte {
	var b Byte = mem.Read(addr)

	*cycles--

	return b
}

// SetBit will set the given bit on the status register
func (b *Byte) SetBit(bit Byte) {
	*b = (*b | bit)
}

// UnsetBit will unset the given bit on the status register
func (b *Byte) UnsetBit(bit Byte) {
	*b &= (^bit)
}

// HasBit checks if the given bit is set on the status register
func (b *Byte) HasBit(bit Byte) bool {
	return (*b & bit) > 0
}
