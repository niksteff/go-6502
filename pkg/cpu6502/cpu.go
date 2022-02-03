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
	N byte = (1 << 7) // Negative flag
	V byte = (1 << 6) // Overflow flag
	U byte = (1 << 5) // unused flag
	B byte = (1 << 4) // Break flag
	D byte = (1 << 3) // Decimal flag (use BCD for arithmetics)
	I byte = (1 << 2) // Interrupt request flag (IRQ disable)
	Z byte = (1 << 1) // Zero flag
	C byte = (1 << 0) // Carry flag
)

// CPU represents a physical 6502 CPU
type CPU struct {
	Bus *Bus  // the address bus
	PC  word  // program counter
	A   byte  // accumulator
	X   byte  // x register
	Y   byte  // y register
	P   byte  // status register, flags are used to set the values
	S   Stack // stack pointer
}

// Reset will reset the CPU into a booting state
func (c *CPU) reset() {
	// set the program counter to the start location
	c.PC = 0xFFFC

	// reset the stackpointer to 0xFF
	c.S = 0xFF
}

// execute will take a byte, decode the instruction in the byte and perform the
// instruction as long as the instruction is a legal OPCODE
func (c *CPU) execute(cyc uint32) {
	for cyc > 0 {
		var ins byte = c.fetchbyte(&cyc, c.PC)
		switch ins {
		case INS_LDA_IM:
			// c2
			c.A = c.fetchbyte(&cyc, c.PC)
			c.setStatusFlags_LDA()
		case INS_LDA_ZP:
			// c2
			var zeroPageAddr byte = c.fetchbyte(&cyc, c.PC)
			// c3
			c.A = c.read(&cyc, word(zeroPageAddr))
			c.setStatusFlags_LDA()
		case INS_LDA_ZPX:
			// c2
			var zeroPageAddr byte = c.fetchbyte(&cyc, c.PC)
			// c3
			var addr byte = c.add(&cyc, zeroPageAddr, c.X)
			// c4
			c.A = c.read(&cyc, word(addr))
		case INS_JMP_ABS:
			// c2
			var adl byte = c.fetchbyte(&cyc, c.PC)
			// c3
			var adh byte = c.fetchbyte(&cyc, c.PC)
			c.PC = ((word(adh) << 8) | word(adl))
		case INS_JSR:
			// c2 fetch adl
			var adl byte = c.fetchbyte(&cyc, c.PC)
			// c3 internal dummy op
			var _ byte = c.Bus.Read(word(c.S))
			cyc--
			// c4 push pch to stack
			// pch -> S
			var t word = c.PC >> 8
			var pch byte = byte(t)
			c.S.Push(&cyc, c.Bus, pch)
			// c5 push pcl to stack
			// pcl -> S
			var pcl byte = byte(c.PC)
			c.S.Push(&cyc, c.Bus, pcl)
			// c6
			c.PC = word(c.Bus.Read(c.PC))<<8 | word(adl)
			cyc--
		case INS_RTS:
			// c2 internal dummy op
			var _ byte = c.read(&cyc, c.PC)
			// c3
			cyc--
			// c4 (s) -> pcl
			var pcl byte = c.S.Pop(&cyc, c.Bus)
			// c5 (s) -> pch
			var pch byte = c.S.Pop(&cyc, c.Bus)
			c.PC = word(pch)<<8 | word(pcl)
			// c6
			c.PC++
			cyc--
		default:
			log.Printf("instruction not handled: %d\n", ins)
		}
	}
}

func (c *CPU) setStatusFlags_LDA() {
	if c.A == 0 {
		// set the zero flag
		SetBit(&c.P, Z)
	}

	if HasBit(&c.A, N) {
		// set the negative flag
		SetBit(&c.P, N)
	}
}

// fetchbyte reads a byte from the given address from the bus
// pc+=1
// cycles-=1
func (c *CPU) fetchbyte(cyc *uint32, addr word) byte {
	var b byte = c.Bus.Read(addr)
	c.PC++
	*cyc--
	
	return b
}

// add will add up two addresses and returns the result
// this can overflow and the resulting address will warp around
// cycles-=1
func (c *CPU) add(cyc *uint32, a byte, b byte) byte {
	var addr byte = a + b
	*cyc--

	return addr
}

// read will read a byte from the bus but does not advance the program counter
// cycles-=1
func (c *CPU) read(cyc *uint32, addr word) byte {
	var b byte = c.Bus.Read(addr)
	*cyc--

	return b
}

// SetBit will set the given bit on the given byte
func SetBit(b *byte, bit byte) {
	*b = (*b | bit)
}

// UnsetBit will unset the given bit on the given byte
func UnsetBit(b *byte, bit byte) {
	*b &= (^bit)
}

// HasBit checks if the given bit is set on the given byte
func HasBit(b *byte, bit byte) bool {
	return (*b & bit) > 0
}
