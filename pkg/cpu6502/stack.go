package cpu6502

type Stack byte

// Push writes the given byte to the next free stack address and sets
// the stack to the next free address in the page
// cycles-=1
func (s *Stack) Push(cyc *uint32, b *Bus, val byte) {
	b.Write(0x0100|word(*s), val)

	*s--
	*cyc--
}

// PushWord will the given word to the stack
// first the high byte of the word will be written
// second the low byte of the word will be written
// cycles-=2
func (s *Stack) PushWord(cyc *uint32, b *Bus, val word) {
	b.Write(0x0100|word(*s), byte(val>>8))
	*s--
	*cyc--

	b.Write(0x0100|word(*s-1), byte(val))
	*s--
	*cyc--
}

// Pop reads the last written byte from the stack pointer and returns the value
// stored at the address on the bus
// cycles-=1
func (s *Stack) Pop(cyc *uint32, b *Bus) byte {
	*s++
	var val byte = b.Read(0x0100 | word(*s))
	*cyc--

	return val
}

// PopWord fetches a word from the stack and returns the value stored on the bus
// for the address of the word
// first the low byte of the word will be read
// second the high byte of the word will be read
// cycles-=2
func (s *Stack) PopWord(cyc *uint32, b *Bus) word {
	// first low byte
	// second high byte
	*s++
	var val word = word(b.Read(0x0100 | word(*s)))
	*cyc--

	*s++
	val |= (word(b.Read(0x0100|word(*s))) << 8)
	*cyc--

	return val
}
