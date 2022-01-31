package cpu6502

// Mem represents the physical RAM
type Mem struct {
	data []Byte
}

// Init will set the size of the physical memory (RAM)
func (m *Mem) init(capacity uint32) {
	m.data = make([]Byte, capacity)
}

func (m *Mem) Access(addr Word) Byte {
	// TODO: create assert to be in memory bounds
	return m.data[addr]
}
