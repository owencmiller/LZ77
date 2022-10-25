package lz77

import "fmt"

func Decompress(compress_data []byte) ([]byte, error) {
	raw_data := make([]byte, 0)
	for i := 0; i < len(compress_data); i += 3 {
		pos := len(raw_data) - int(compress_data[i])
		length := pos + int(compress_data[i+1])
		if length < 0 || pos < 0 || pos+length > len(compress_data) {
			return nil, fmt.Errorf("lz77: out of bounds read. i %d, pos %d, length %d", i, pos, length)
		}
		raw_data = append(raw_data, raw_data[pos:length]...)
		raw_data = append(raw_data, compress_data[i+2])
	}
	return raw_data, nil
}
