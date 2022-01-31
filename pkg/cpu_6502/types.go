package cpu6502

// every flag has a bit inside the status register this ordered from bit 7 to
// bit 0. We do not have single bits in golang so to save space we will create
// these constants which will each contain the bit to set for the flag. If the
// bit will not be set we can just not add the value to the status register
const (
	N  uint8 = 7 // Negative flag
	V  uint8 = 6 // Overflow flag
	_I uint8 = 5 // ignored flag
	B  uint8 = 4 // Break flag
	D  uint8 = 3 // Decimal flag (use BCD for arithmetics)
	I  uint8 = 2 // Interrupt flag (IRQ disable)
	Z  uint8 = 1 // Zero flag
	C  uint8 = 0 // Carry flag

	MEM_CAP_apple2 uint32 = 1024 * 64 // set the max memory to 64KiB
)

// Byte represents an unsigned 8 bit integer
type Byte uint8

// Word represents two Bytes
type Word uint16
