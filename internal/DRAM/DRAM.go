package DRAM

import (
	"errors"
	"fmt"
)

const Size uint64 = 1 * 1024 * 1024 // 1 MiB
const BaseAddr uint64 = 0x80000000  // Lower addresses are reserved for memory-mapped I/O

// DRAM is an array of (unsigned) bytes that represents memory
type DRAM struct {
	content [Size]uint8
}

// Load takes the pointer to the dram to be read from, the address of the data to be read, and the size of the data to be read, which might be 8, 16, 32 or 64 bits, as per the instruction (LB, LH, LW, and LD respectively).
func Load(memPtr *DRAM, addr uint64, size uint64) (uint64, error) {
	if (addr - BaseAddr) > Size {
		return 1, errors.New("invalid memory address")
	}
	switch size {
	case 64:
		return uint64(memPtr.content[addr-BaseAddr]) |
			(uint64(memPtr.content[addr-BaseAddr+1]) << 8) |
			(uint64(memPtr.content[addr-BaseAddr+2]) << 16) |
			(uint64(memPtr.content[addr-BaseAddr+3]) << 24) |
			(uint64(memPtr.content[addr-BaseAddr+4]) << 32) |
			(uint64(memPtr.content[addr-BaseAddr+5]) << 40) |
			(uint64(memPtr.content[addr-BaseAddr+6]) << 48) |
			(uint64(memPtr.content[addr-BaseAddr+7]) << 56), nil
	case 32:
		return uint64(memPtr.content[addr-BaseAddr]) |
			(uint64(memPtr.content[addr-BaseAddr+1]) << 8) |
			(uint64(memPtr.content[addr-BaseAddr+2]) << 16) |
			(uint64(memPtr.content[addr-BaseAddr+3]) << 24), nil
	case 16:
		return uint64(memPtr.content[addr-BaseAddr]) |
			(uint64(memPtr.content[addr-BaseAddr+1]) << 8), nil
	case 8:
		return uint64(memPtr.content[addr-BaseAddr]), nil
	default:
		return 1, fmt.Errorf("Load: invalid size: %d", size)
	}
}

// Store stores
func Store(memPtr *DRAM, addr uint64, size uint64, value uint64) error {
	if (addr-BaseAddr) > Size || (addr-BaseAddr) < 0 {
		return errors.New("invalid memory address")
	}
	switch size {
	case 64:
		memPtr.content[addr-BaseAddr+0] = uint8(value)
		memPtr.content[addr-BaseAddr+1] = uint8((value >> 8))
		memPtr.content[addr-BaseAddr+2] = uint8((value >> 16))
		memPtr.content[addr-BaseAddr+3] = uint8((value >> 24))
		memPtr.content[addr-BaseAddr+4] = uint8((value >> 32))
		memPtr.content[addr-BaseAddr+5] = uint8((value >> 40))
		memPtr.content[addr-BaseAddr+6] = uint8((value >> 48))
		memPtr.content[addr-BaseAddr+7] = uint8((value >> 56))
	case 32:
		memPtr.content[addr-BaseAddr+0] = uint8(value)
		memPtr.content[addr-BaseAddr+1] = uint8((value >> 8))
		memPtr.content[addr-BaseAddr+2] = uint8((value >> 16))
		memPtr.content[addr-BaseAddr+3] = uint8((value >> 24))
	case 16:
		memPtr.content[addr-BaseAddr+0] = uint8(value)
		memPtr.content[addr-BaseAddr+1] = uint8((value >> 8))
	case 8:
		memPtr.content[addr-BaseAddr+0] = uint8(value)
	default:
		return fmt.Errorf("Load: invalid size: %d", size)
	}
	return nil
}
