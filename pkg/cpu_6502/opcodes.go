package cpu6502

const (
	INS_LDA_IM  Byte = 0xA9 // Load acummulator immediately, 2 Bytes, 2 Cycles
	INS_LDA_ZP  Byte = 0xA5 // Load accumulator from zeroPage, 2 Bytes, 3 Cycles
	INS_LDA_ZPX Byte = 0xB5 // Load accumulator from zeroPage X, 2 Bytes, 4 Cycles
	INS_JMP_ABS Byte = 0x4C // Jump to, 3 Bytes, 6 Cycles
	INS_JMP_IND Byte = 0x6C // Jump to indirect, // TODO
	INS_JSR     Byte = 0x20 // Jump to subroutine, 3 Bytes, 6 Cycles
	INS_RTS     Byte = 0x60 // Return from subroutine, 1 Byte, 6 Cycles
)
