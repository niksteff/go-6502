package cpu6502

// ROM is a read only memory data structure represented as a slice of bytes
type ROM struct {
	data []byte
}

// Read returns the byte stored at the given address in the rom
func (r *ROM) Read(addr word) byte {
	return r.data[addr]
}
