package cpu6502

const (
	INS_LDA_IM  Byte = 0xA9 // 2 Bytes, 2 Cycles
	INS_LDA_ZP  Byte = 0xA5 // 2 Bytes, 3 Cycles
	INS_LDA_ZPX Byte = 0xB5 // 2 Bytes, 4 Cycles
	INS_JMP_ABS Byte = 0x4C // 3 Bytes, 6 Cycles
	INS_JMP_IND Byte = 0x6C // TODO
	INS_JSR_ABS Byte = 0x20 // 3 Bytes, 6 Cycles
)
