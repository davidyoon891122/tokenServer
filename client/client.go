package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	var IP string = "10.131.150.171"
	var PORT string = ":13302"

	conn, err := net.Dial("tcp", IP+PORT)

	defer conn.Close()
	if err != nil {
		panic(err)
	}
	data := packMsg()

	for {
		var line string
		conn.Write(data)
		fmt.Scanln(&line)
	}
}

func packMsg() []byte {
	var bytesBuffer bytes.Buffer

	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(16))
	bytesBuffer.Write(buf)
	buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(0x7FF3))
	bytesBuffer.Write(buf)
	buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(0))
	bytesBuffer.Write(buf)
	buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(1050))
	bytesBuffer.Write(buf)
	bytesBuffer.WriteByte(byte(0))
	bytesBuffer.WriteByte(byte(0))
	buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(0))
	bytesBuffer.Write(buf)

	bytesBuffer.Write([]byte("eR94gyvABa8:APA91bHpDrn3xN3CqWPNfybP3ztK-_QJUi0lksxXcPcaaoCyKCD96CIP3a7SzlYb68a2yy9a2TLCK8SVFOcNvSGkceN3PJBZeXQXF15oekzI_whVMmi5Sk-Xtn6tZ3XrvlebhbF1oTlbhtsadmin"))

	return bytesBuffer.Bytes()
}
