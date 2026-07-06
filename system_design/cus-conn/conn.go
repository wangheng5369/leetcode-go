package cus_conn

import (
	"errors"
	"io"
	"net"
)

const (
	FrameTypeHeader byte   = 0x01
	FrameTypeData   byte   = 0x02
	FrameTypeEnd    byte   = 0x04
	HeaderLen              = 9
	MaxMakeSliceLen uint64 = 1 << 12
)

// Conn 是你需要实现的一种连接类型，它支持下面描述的若干接口；
// 为了实现这些接口，你需要设计一个基于 TCP 的简单协议；
type Conn struct {
	tcpConn net.Conn
}

// Send 传入一个 key 表示发送者将要传输的数据对应的标识；
// 返回 writer 可供发送者分多次写入大量该 key 对应的数据；
// 当发送者已将该 key 对应的所有数据写入后，调用 writer.Close 告知接收者该 key 的数据已经完全写入；
func (conn *Conn) Send(key string) (writer io.WriteCloser, err error) {
	header := &Header{
		Type: FrameTypeHeader,
		Len:  uint64(len(key)),
	}
	data := append(header.Marshal(), []byte(key)...)
	_, err = conn.tcpConn.Write(data)
	if err != nil {
		return nil, err
	}
	writer = &curWriter{tcpConn: conn.tcpConn}
	return
}

// Receive 返回一个 key 表示接收者将要接收到的数据对应的标识；
// 返回的 reader 可供接收者多次读取该 key 对应的数据；
// 当 reader 返回 io.EOF 错误时，表示接收者已经完整接收该 key 对应的数据；
func (conn *Conn) Receive() (key string, reader io.Reader, err error) {
	headerData := make([]byte, HeaderLen)
	_, err = conn.tcpConn.Read(headerData)
	if err != nil {
		return "", nil, err
	}
	header := &Header{}
	header.Unmarshal(headerData)
	if header.Type != FrameTypeHeader {
		return "", nil, errors.New("invalid header type")
	}
	keyData := make([]byte, header.Len)
	_, err = conn.tcpConn.Read(keyData)
	if err != nil {
		return "", nil, err
	}
	key = string(headerData)
	reader = &cusReader{tcpConn: conn.tcpConn, buffer: make([]byte, 0)}
	return
}

// Close 关闭你实现的连接对象及其底层的 TCP 连接
func (conn *Conn) Close() {
	conn.tcpConn.Close()
}

// NewConn 从一个 TCP 连接得到一个你实现的连接对象
func NewConn(conn net.Conn) *Conn {
	return &Conn{tcpConn: conn}
}
