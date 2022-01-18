package lz77

import (
	"math"
)

const SEARCH_BUFFER = 4096
const LOOKAHEAD_BUFFER = 12

func CompressString(raw_data string) []byte {
	return Compress([]byte(raw_data))
}

func Compress(raw_data []byte) []byte {
	compress_data := make([]byte, 0)
	for pos := 0; pos < len(raw_data); pos++ {

		searchPos := int(math.Max(float64(pos-SEARCH_BUFFER), 0))
		lookPos := int(math.Min(float64(pos+LOOKAHEAD_BUFFER), float64(len(raw_data))))

		offset, length, char := find(raw_data[searchPos:pos], raw_data[pos:lookPos])
		pos += int(length)

		compress_data = append(compress_data, offset, length, char)
	}
	return compress_data
}

func find(search []byte, look_ahead []byte) (byte, byte, byte) {
	offset, length := 0, 0
	char := look_ahead[0]

	for i, c := range search {
		if c == look_ahead[0] {
			for j, let := range look_ahead {
				if i+j < len(search) && let == search[i+j] {
					if j >= length {
						offset = len(search) - i
						length = j + 1
						if length >= len(look_ahead) {
							char = 0
						} else {
							char = look_ahead[j+1]
						}
					}
				} else {
					break
				}
			}
		}
	}
	return byte(offset), byte(length), char
}
