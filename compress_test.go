package lz77

import "testing"

func TestCompressString(t *testing.T) {
	data := "cabracadabrarrarrad"
	compressed := CompressString(data)
	println("Compressed Data:", compressed)
}

func TestFind(t *testing.T) {
	search := []byte("")
	look_ahead := []byte("is a t")
	println(find(search, look_ahead))
}

func TestCompressDecompress(t *testing.T) {
	data := "cabracadabrarrarrad"
	compressed := CompressString(data)
	new_data := Decompress(compressed)
	if string(new_data) != data {
		t.Error("Mismatch:", data, new_data)
	}
}
