package tools


import (
    "fmt"
    "../headerStruct"
    "encoding/binary"
)


//Global Variables
var recvData []byte
var header headerStruct.Header
var index int
//var body bodyStruct.Body

func Parse(data []byte) {
    fmt.Println("Parser is initializing..")
    recvData = data


    //ParseHeader
    parseHeader()

}


func parseHeader() {
    index = 0
    header.Length = readInt()
    header.Process = readShort()
    header.Service = readShort()
    header.Window = readInt()
    header.Control = readByte()
    header.Flag = readByte()
    header.Reserve = readInt()
    fmt.Println("Header : ", header)
}




//func parseBody() {
//    body.Token = readFixLenStr(135)
//    body.UserID = readFixLenStr(15)

//}


func readInt() int {
    ret := binary.BigEndian.Uint32(recvData[index:])
    fmt.Println("readInt")
    index += 4
    return int(ret)
}


func readShort() int16{
    fmt.Println("readShort")
    ret := binary.BigEndian.Uint16(recvData[index:])
    index += 2
    return int16(ret)
}


func readByte() byte {
    fmt.Println("readByte")
    ret := recvData[index:index+1]
    index += 1
    return ret[0]
}






