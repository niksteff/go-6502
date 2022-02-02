package main

import (
	"log"

	cpu6502 "github.com/niksteff/go-6502/pkg/cpu_6502"
)

func main() {
	mem := cpu6502.Mem{}
	cpu := cpu6502.CPU{}
	
	cpu.Reset(&mem)
	
	// write a little program in memory
	mem.Write(0xFFFC, cpu6502.INS_JSR_ABS)
	mem.Write(0xFFFD, 0x0A)
	mem.Write(0xFFFE, 0x0B)
	
	mem.Write(0x0B0A, cpu6502.INS_LDA_IM)
	mem.Write(0x0B0B, 0x42)
	
	cpu.Execute(&mem, 8)

	log.Println("exiting ...")
}
