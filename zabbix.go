package mlib

import (
	"bytes"
	"encoding/binary"
	"net"
	"time"
)

const (
	HeaderString     = "ZBXD"
	HeaderLength     = len(HeaderString)
	HeaderVersion    = uint8(1)
	DataLengthOffset = int64(HeaderLength + 1)
	DataLengthSize   = int64(8)
	DataOffset       = int64(DataLengthOffset + DataLengthSize)
	ErrorMessage     = "ZBX_NOTSUPPORTED"
)

var (
	ErrorMessageBytes = []byte(ErrorMessage)
	Terminator        = []byte("\n")
	HeaderBytes       = []byte(HeaderString)
)

type ConnReader interface {
	Read(b []byte) (n int, err error)
}

func Data2Packet(data []byte) []byte {
	buf := new(bytes.Buffer)
	buf.Write(HeaderBytes)
	binary.Write(buf, binary.LittleEndian, HeaderVersion)
	binary.Write(buf, binary.LittleEndian, int64(len(data)))
	buf.Write(data)
	return buf.Bytes()
}

func Stream2Data(conn ConnReader) (rdata []byte, err error) {
	// read header "ZBXD\x01"
	head := make([]byte, DataLengthOffset)
	_, err = conn.Read(head)
	if err != nil {
		return
	}
	if bytes.Equal(head[0:HeaderLength], HeaderBytes) && head[HeaderLength] == byte(HeaderVersion) {
		rdata, err = parseBinary(conn)
	} else {
		rdata, err = parseText(conn, head)
	}
	return
}

func parseBinary(conn ConnReader) (rdata []byte, err error) {
	// read data length
	var dataLength int64
	err = binary.Read(conn, binary.LittleEndian, &dataLength)
	if err != nil {
		return
	}
	// read data body
	buf := make([]byte, 1024)
	data := new(bytes.Buffer)
	total := 0
	size := 0
	for total < int(dataLength) {
		size, err = conn.Read(buf)
		if err != nil {
			return
		}
		if size == 0 {
			break
		}
		total = total + size
		data.Write(buf[0:size])
	}
	rdata = data.Bytes()
	return
}

func parseText(conn ConnReader, head []byte) (rdata []byte, err error) {
	data := new(bytes.Buffer)
	data.Write(head)
	buf := make([]byte, 1024)
	size := 0
	for {
		// read data while "\n" found
		size, err = conn.Read(buf)
		if err != nil {
			return
		}
		if size == 0 {
			break
		}
		i := bytes.Index(buf[0:size], Terminator)
		if i == -1 {
			// terminator not found
			data.Write(buf[0:size])
			continue
		}
		// terminator found
		data.Write(buf[0 : i+1])
		break
	}
	rdata = data.Bytes()
	return
}

// ZabbixGet is like command zabbix_get
func ZabbixGet(addr string, key string, timeout time.Duration) (value string, err error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(timeout))

	msg := Data2Packet([]byte(key))
	_, err = conn.Write(msg)
	if err != nil {
		return
	}
	_value, err := Stream2Data(conn)
	return string(_value), err
}
