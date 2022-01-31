package cpu6502

// every flag has a bit inside the status register this ordered from bit 7 to
// bit 0. We do not have single bits in golang so to save space we will create
// these constants which will each contain the bit to set for the flag. If the
// bit will not be set we can just not add the value to the status register
const (
	N  uint8 = 0x80 // Negative flag
	V  uint8 = 0x40 // Overflow flag
	_I uint8 = 0x20 // ignored flag
	B  uint8 = 0x10 // Break flag
	D  uint8 = 0x08 // Decimal flag (use BCD for arithmetics)
	I  uint8 = 0x04 // Interrupt flag (IRQ disable)
	Z  uint8 = 0x02 // Zero flag
	C  uint8 = 0x01 // Carry flag

	MEM_CAP_apple2 uint32 = 1024*64 // set the max memory to 64KiB
)

// Byte represents an unsigned 8 bit integer
type Byte uint8

// Word represents two Bytes
type Word uint16