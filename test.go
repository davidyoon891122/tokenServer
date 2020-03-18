package main

import (
    "fmt"
    "./tools"
    "encoding/binary"
    "bytes"
)


func main(){
    fmt.Println("Test is starting.")
    data := makePack()
    fmt.Println(data)
    tools.Parse(data)

}


func makePack() []byte {
    var byteBuffer bytes.Buffer
    buf := make([]byte, 4)
    binary.BigEndian.PutUint32(buf, uint32(16))
    byteBuffer.Write(buf)
    buf = make([]byte, 2)
    binary.BigEndian.PutUint16(buf, uint16(0xEE))
    byteBuffer.Write(buf)
    buf = make([]byte, 2)
    binary.BigEndian.PutUint16(buf, uint16(0xFF))
    byteBuffer.Write(buf)
    buf = make([]byte, 4)
    binary.BigEndian.PutUint32(buf, uint32(3))
    byteBuffer.Write(buf)
    byteBuffer.WriteByte(1)
    byteBuffer.WriteByte(2)
    buf = make([]byte, 4)
    binary.BigEndian.PutUint32(buf, uint32(20))
    byteBuffer.Write(buf)
    return byteBuffer.Bytes()
}
