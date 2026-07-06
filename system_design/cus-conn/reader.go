package cus_conn

import (
	"io"
	"net"
)

type cusReader struct { // 3 usages
	tcpConn net.Conn
	buffer  []byte
	isEnd   bool
}

func (r *cusReader) writBuffer(dataLen uint64) error { // 1 usage
	sliceLen := dataLen
	if dataLen > MaxMakeSliceLen {
		sliceLen = MaxMakeSliceLen
	}
	var readTotal uint64 = 0
	for readTotal < dataLen {
		tmp := make([]byte, sliceLen)
		n, err := r.tcpConn.Read(tmp)
		if err != nil {
			return err
		}
		r.buffer = append(r.buffer, tmp[:n]...)
		readTotal += uint64(n)
		if dataLen-readTotal < sliceLen {
			sliceLen = dataLen - readTotal
		}
	}
	return nil
}

func (r *cusReader) Read(p []byte) (n int, err error) {
again:
	// 1. 如果本地缓冲区 buffer 中的数据足够填满 p，或者够这次读取，则直接从缓冲区复制并返回
	if len(p) <= len(r.buffer) {
		copy(p, r.buffer[:len(p)])
		r.buffer = r.buffer[len(p):]
		return len(p), nil
	}

	// 2. 如果已经标记为结束状态 (isEnd)
	if r.isEnd {
		if len(r.buffer) == 0 {
			// 所有数据均已读取完成
			return 0, io.EOF
		}
		// 如果缓冲区还有残余数据，全部吐给 p
		copy(p, r.buffer)
		n := len(r.buffer)
		r.buffer = []byte{}
		return n, nil
	}

	// 3. 缓冲区不够，且未结束，开始从底层的 TCP 连接中读取报文头 (Header)
	headerData := make([]byte, HeaderLen)
	n, err = r.tcpConn.Read(headerData)
	if err == io.EOF {
		r.isEnd = true
		goto again
	}
	if err != nil {
		return 0, err
	}

	// 4. 解析报文头
	header := &Header{}
	header.UnMarshal(headerData)

	// 5. 如果报文类型是结束帧，标记 isEnd 并跳转重试（去处理缓冲区残余数据或返回 EOF）
	if header.Type == FrameTypeEnd {
		r.isEnd = true
		goto again
	}

	// 6. 如果不是结束帧，调用 writBuffer 将指定长度的数据读入本地 buffer 中
	err = r.writBuffer(header.Len)
	if err != nil {
		return 0, err
	}

	// 7. 数据存入 buffer 后，重新跳转到 again 处，将数据复制给入参 p
	goto again
}
