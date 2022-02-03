package cpu6502

const (
	INS_LDA_IM  byte = 0xA9 // Load acummulator immediately, 2 bytes, 2 Cycles
	INS_LDA_ZP  byte = 0xA5 // Load accumulator from zeroPage, 2 bytes, 3 Cycles
	INS_LDA_ZPX byte = 0xB5 // Load accumulator from zeroPage X, 2 bytes, 4 Cycles
	INS_JMP_ABS byte = 0x4C // Jump to, 3 bytes, 6 Cycles
	INS_JMP_IND byte = 0x6C // Jump to indirect, // TODO
	INS_JSR     byte = 0x20 // Jump to subroutine, 3 bytes, 6 Cycles
	INS_RTS     byte = 0x60 // Return from subroutine, 1 byte, 6 Cycles
)
