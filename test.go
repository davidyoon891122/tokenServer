package main

import (
    "fmt"
    "./tools"
    "encoding/binary"
    "bytes"
    "time"
    "math/rand"
)


func main(){
    fmt.Println("Test is starting.")
    data := makePack()
    fmt.Println(data)
    tools.Parse(data)

}


func makePack() []byte {
    var byteBuffer bytes.Buffer
    //pack header
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


    //pack body
    buf = make([]byte, 135)
    buf = []byte(randString(135))
    byteBuffer.Write(buf)

    buf = make([]byte, 15)
    buf = []byte("htsadmin")
    byteBuffer.Write(buf)
    return byteBuffer.Bytes()
}


func randString(n int) string {
    buf := make([]rune, n)
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ")
    now := time.Now().Second()
    rand.Seed(int64(now))

    for i := range buf {
        buf[i] = letterRunes[rand.Intn(len(letterRunes))]
    }

    return string(buf)
}



