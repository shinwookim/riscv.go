package DRAM

import (
	"math/rand"
	"testing"
)

func randRange(min, max uint64) uint64 {
	return uint64(rand.Int63n(int64(max-min)) + int64(min))
}

func TestDRAM(t *testing.T) {
	var testMemory DRAM
	err := Store(&testMemory, BaseAddr, 8, 15)
	if err != nil {
		t.Fatalf("Error while storing: %v", err)
	}
	data, err := Load(&testMemory, BaseAddr, 8)
	if err != nil {
		t.Fatalf("Error while loading: %v", err)
	}
	if data != 15 {
		t.Errorf("Wrong data: %d", data)
	}
}

func TestDRAMSingleByte(t *testing.T) {
	var testMemory DRAM
	// Generate Random Address-Data Pairs
	m := make(map[uint64]uint64)
	for i := BaseAddr; i < BaseAddr+Size; i++ {
		m[i] = rand.Uint64()
		Store(&testMemory, i, 8, m[i])
	}
	for i := BaseAddr + Size - 1; i >= BaseAddr; i-- {
		data, err := Load(&testMemory, i, 8)
		if err != nil {
			t.Fatalf("Error while loading: %v", err)
		}
		if data != uint64(uint8(m[i])) {
			t.Errorf("Wrong data: %d\t Expected %d", data, uint64(uint8(m[i])))
		}
	}

}

func TestDRAMTwoBytes(t *testing.T) {
	var testMemory DRAM
	// Generate Random Address-Data Pairs
	m := make(map[uint64]uint64)
	var topAddr uint64 = BaseAddr
	for ; topAddr < BaseAddr+Size; topAddr += 2 {
		m[topAddr] = rand.Uint64()
		Store(&testMemory, topAddr, 16, m[topAddr])
	}
	for i := topAddr - 2; i >= BaseAddr; i -= 2 {
		data, err := Load(&testMemory, i, 16)
		if err != nil {
			t.Fatalf("Error while loading: %v", err)
		}
		if data != uint64(uint16(m[i])) {
			t.Errorf("Wrong data: %d\t Expected %d", data, uint64(uint16(m[i])))
		}
	}
}
func TestDRAMFourBytes(t *testing.T) {
	var testMemory DRAM
	// Generate Random Address-Data Pairs
	m := make(map[uint64]uint64)
	var topAddr uint64 = BaseAddr
	for ; topAddr < BaseAddr+Size; topAddr += 4 {
		m[topAddr] = rand.Uint64()
		Store(&testMemory, topAddr, 32, m[topAddr])
	}
	for i := topAddr - 4; i >= BaseAddr; i -= 4 {
		data, err := Load(&testMemory, i, 32)
		if err != nil {
			t.Fatalf("Error while loading: %v", err)
		}
		if data != uint64(uint32(m[i])) {
			t.Errorf("Wrong data: %d\t Expected %d", data, uint64(uint32(m[i])))
		}
	}
}
func TestDRAMEightBytes(t *testing.T) {
	var testMemory DRAM
	// Generate Random Address-Data Pairs
	m := make(map[uint64]uint64)
	var topAddr uint64 = BaseAddr
	for ; topAddr < BaseAddr+Size; topAddr += 8 {
		m[topAddr] = rand.Uint64()
		Store(&testMemory, topAddr, 64, m[topAddr])
	}
	for i := topAddr - 8; i >= BaseAddr; i -= 8 {
		data, err := Load(&testMemory, i, 64)
		if err != nil {
			t.Fatalf("Error while loading: %v", err)
		}
		if data != uint64(m[i]) {
			t.Errorf("Wrong data: %d\t Expected %d at address %d", data, m[i], i)
		}
	}
}

func TestDRAMRandBytes(t *testing.T) {}
