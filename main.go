package main

import (
	"log"

	cpu6502 "github.com/niksteff/go-6502/pkg/cpu6502"
)

func main() {
	bus := cpu6502.New()
	
	// program start
	bus.Write(0xFFFC, cpu6502.INS_JSR)
	bus.Write(0xFFFD, 0x0A)
	bus.Write(0xFFFE, 0x0B)
	
	bus.Write(0xFFFF, cpu6502.INS_LDA_IM)
	bus.Write(0x0, 0xA)

	bus.Write(0x0B0A, cpu6502.INS_LDA_IM)
	bus.Write(0x0B0B, 0x42)
	bus.Write(0x0B0C, cpu6502.INS_RTS)
	// program end
	
	bus.Execute(16)

	log.Println("exiting ...")
}
