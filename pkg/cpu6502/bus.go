package cpu6502

// word represents an unsigned 16bit integer this is a convenience type that
// enables us to write word instead of uint16
type word uint16

// Bus is an internal communication bus that connects all parts of the
// microprocessor
type Bus struct {
	cpu CPU
	ram RAM
	rom ROM
}

// New returns a new instance of the microprocessors bus which can be used to
// read and write from and to it
func New() Bus {
	var b Bus
	b = Bus{
		cpu: CPU{
			bus: &b,
		},
		ram: RAM{
			data: make([]byte, 1024*64),
		},
		rom: ROM{
			data: make([]byte, 1024*64 ), // TODO: correct the rom size
		},
	}

	b.reset()

	return b
}

// reset the microprocessor
func (b *Bus) reset() {
	// create 64KiB of memory
	b.ram.data = make([]byte, 1024*64)

	// reset the cpu to read the first instruction
	b.cpu.reset()

	// TODO: call execute with the correct amount of instructions
	// TODO: use the ROM to load the boot sequence
}

// Execute the given number of cycles
func (b *Bus) Execute(cyc uint32) {
	b.cpu.execute(cyc)
}

// Read 1 byte from the given address on the bus
func (b *Bus) Read(addr word) byte {
	// TODO: create assert to be in memory bounds
	var d byte = b.ram.data[addr]
	return d
}

// Write 1 byte to the bus at the given address
func (b *Bus) Write(addr word, val byte) {
	// TODO: create assert to be in memory bounds
	b.ram.data[addr] = val
}
