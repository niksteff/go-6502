package cpu6502

import (
	"bytes"
	"io"
	"log"
)

// IODevice is a device capable of sending and receiving data like device
// connected to an I/O port
type IODevice interface {
	io.ReadWriter
}

// IDevice is capable of writing data to the bus
type IDevice interface {
	io.Writer
}

// ODevice is capable of reading data from the bus
type ODevice interface {
	io.Reader
}

// DiskII is a floppy disk peripheral input device
type DiskII struct {
	// io.SectionReader
	bytes.Buffer
}

// Load reads in a disk from a byte buffer
func (d *DiskII) Load(b io.Reader) {
	// TODO: read a file from disk into a buffer and write it into the data
	// structure as a byte slice
	len, err := d.ReadFrom(b)
	if err != nil {
		log.Println(err)
	}

	log.Printf("read %d bytes from diskII", len)
}

func (d *DiskII) Read() byte {
	// TODO: read back one byte at a time from the stored data
	return 0
}
