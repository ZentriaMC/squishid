package squishid

import (
	"errors"
	"strings"
)

var RestoreError = errors.New("RestoreError")

const squishTable = "abcdefghijkmopqrstuwxyz123456789" // l0nv

var repeatSquishTable = map[uint8][]uint8{
	3:  []uint8("l"),
	4:  []uint8("0"),
	5:  []uint8("v"),
	6:  []uint8("n"),
	7:  []uint8("l0"),
	8:  []uint8("lv"),
	9:  []uint8("ln"),
	10: []uint8("0n"),
	11: []uint8("vn"),
	12: []uint8("nn"),
}

var restoreTable = map[uint8]uint8{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
	'i': 8,
	'j': 9,
	'k': 10,
	'm': 11,
	'o': 12,
	'p': 13,
	'q': 14,
	'r': 15,
	's': 16,
	't': 17,
	'u': 18,
	'w': 19,
	'x': 20,
	'y': 21,
	'z': 22,
	'1': 23,
	'2': 24,
	'3': 25,
	'4': 26,
	'5': 27,
	'6': 28,
	'7': 29,
	'8': 30,
	'9': 31,
}

var repeatRestoreTable = map[uint8]uint8{
	'l': 3,
	'0': 4,
	'v': 5,
	'n': 6,
}

func Squish(id uint64) string {
	var buffer strings.Builder
	var repeatCount uint8 = 0
	var lastChunk uint64 = 0

	write := func() {
		char := squishTable[lastChunk]
		if repeatCount == 2 {
			buffer.WriteByte(char)
		} else if repeatCount > 2 {
			buffer.Write(repeatSquishTable[repeatCount])
		}
		buffer.WriteByte(char)
		repeatCount = 0
	}

	for id > 0 {
		chunk := id & 0b11111
		if (chunk != lastChunk && repeatCount > 0) || repeatCount == 12 {
			write()
		}
		lastChunk = chunk
		repeatCount += 1
		id = id >> 5
	}
	write()

	return buffer.String()
}

func Restore(id string) (uint64, error) {
	var result uint64 = 0
	shift := 0
	const maxShift = 13 * 5
	var repeatCount uint8 = 0

	for i := 0; i < len(id); i++ {
		char := id[i]

		isUpperCase := char >= 65 && char <= 90
		if isUpperCase {
			char += 32
		}

		if chunk, exists := restoreTable[char]; exists {
			for done := false; !done; done = repeatCount == 0 {
				result |= uint64(chunk) << shift
				shift += 5
				if repeatCount > 0 {
					repeatCount--
				}
			}
		} else if count, exists := repeatRestoreTable[char]; exists {
			repeatCount += count
		} else {
			return 0, RestoreError
		}

		if shift+int(repeatCount)*5 > maxShift {
			return 0, RestoreError
		}
	}

	return result, nil
}
