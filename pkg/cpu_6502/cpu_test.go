package cpu6502_test

import (
	"log"
	"testing"

	cpu6502 "github.com/niksteff/go-6502/pkg/cpu_6502"
)

const (
	DEBUG bool = false
)

func Debug(s string, args ...interface{}) {
	if DEBUG {
		log.Printf(s, args...)
	}
}

func TestSetBit(t *testing.T) {
	var want cpu6502.Byte = 0b_1000_0010

	var sr cpu6502.Byte = 0b_0000_0000
	sr.SetBit(cpu6502.Z | cpu6502.N)

	if want != sr {
		t.Errorf("expected the result of the bitwise or connection to be %08b but got %08b", want, sr)
	}

	Debug("TestBitOr: %08b", sr)
}

func TestUnsetBit(t *testing.T) {
	var want cpu6502.Byte = 0b_0001_0000

	var sr cpu6502.Byte = 0b_0001_0010
	sr.UnsetBit(cpu6502.Z)

	if want != sr {
		t.Errorf("expected the result of the bitwise or connection to be %08b but got %08b", want, sr)
	}

	Debug("TestBitOr: %08b", sr)
}

func TestHasFlag(t *testing.T) {
	var want bool = true

	var sr cpu6502.Byte = 0b_0000_0010

	var isSet bool
	if sr.HasBit(cpu6502.Z) {
		isSet = true
	}

	if want != isSet {
		t.Errorf("expected the result of the bitwise or connection to be %t but got %t", want, isSet)
	}

	Debug("TestBitAnd: %t", isSet)
}

func TestNotHasFlag(t *testing.T) {
	var want bool = false

	var sr cpu6502.Byte = 0b_0000_0000

	var isSet bool
	if sr.HasBit(cpu6502.N) {
		isSet = true
	}

	if want != isSet {
		t.Errorf("expected the result of the bitwise or connection to be %t but got %t", want, isSet)
	}

	Debug("TestBitAnd: %t", isSet)
}