package lz77

func Decompress(compress_data []byte) []byte {
	raw_data := make([]byte, 0)
	for i := 0; i < len(compress_data); i += 3 {
		pos := len(raw_data) - int(compress_data[i])
		length := pos + int(compress_data[i+1])
		raw_data = append(raw_data, raw_data[pos:length]...)
		raw_data = append(raw_data, compress_data[i+2])
	}
	return raw_data
}
