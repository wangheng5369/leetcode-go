package cus_conn

import "net"

type curWriter struct {
	tcpConn net.Conn
}

func (w *curWriter) Write(b []byte) (int, error) {
	header := &Header{
		Type: FrameTypeData,
		Len:  uint64(len(b)),
	}
	data := append(header.Marshal(), b...)
	n, err := w.tcpConn.Write(data)
	n -= HeaderLen
	return n, err
}

func (w *curWriter) Close() error {
	header := &Header{
		Type: FrameTypeEnd,
		Len:  0,
	}
	data := header.Marshal()
	_, err := w.tcpConn.Write(data)
	return err
}
