package cus_conn

import "encoding/binary"

type Header struct {
	Type byte
	Len  uint64
}

func (header *Header) Marshal() []byte {
	b := []byte{header.Type}
	lenB := make([]byte, 8)
	binary.BigEndian.PutUint64(lenB, header.Len)
	return append(b, lenB...)
}

func (header *Header) UnMarshal(data []byte) {
	header.Type = data[0]
	lenB := data[1:]
	header.Len = binary.BigEndian.Uint64(lenB)
}
