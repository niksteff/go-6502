package main

import (
	"log"

	cpu6502 "github.com/niksteff/go-6502/pkg/cpu_6502"
)

func main() {
	mem := cpu6502.Mem{}
	cpu := cpu6502.CPU{}
	
	cpu.Reset(&mem)

	cpu.Execute(&mem, 2)

	log.Println("exiting ...")
}
