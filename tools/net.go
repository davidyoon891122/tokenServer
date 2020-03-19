package tools


import (
    "fmt"
    "../headerStruct"
    "../bodyStruct"
    "encoding/binary"
    "strconv"
)


//Global Variables
var recvData []byte
var header headerStruct.Header
var index int
var body bodyStruct.Body

func Parse(data []byte) bodyStruct.Body {
    fmt.Println("Parser is initializing..")
    recvData = data


    //Parse Header
    parseHeader()
    //Parse Body
    parseBody()

    printHeader()
    //data Processing.

    return body
}


func parseHeader() {
    index = 0
    header.Length = readInt()
    header.Process = readShort()
    header.Service = readShort()
    header.Window = readInt()
    header.Control = readByte()
    header.Flag = readByte()
    header.Reserve = readShort()
    fmt.Println("Header : ", header)

    service := getService()
    fmt.Printf("Servcie number : %s\n", service)
}



func printHeader() {
    fmt.Println("Length : ", header.Length)
    fmt.Println("Process : ", header.Process)
    fmt.Println("Service : ", header.Service)
    fmt.Println("Window : ", header.Window)
    fmt.Println("Control : ", header.Control)
    fmt.Println("Flag : ", header.Flag)
    fmt.Println("Reserve : ", header.Reserve)
}



func parseBody() {
    body.Token = readFixLenStr(bodyStruct.TokenSize)
    body.UserID = readFixLenStr(bodyStruct.UserIDSize)

}


func readInt() int {
    ret := binary.BigEndian.Uint32(recvData[index:])
    index += 4
    return int(ret)
}


func readShort() int16{
    ret := binary.BigEndian.Uint16(recvData[index:])
    index += 2
    return int16(ret)
}


func readByte() byte {
    ret := recvData[index:index+1]
    index += 1
    return ret[0]
}


func readFixLenStr(length int) string{
    ret := recvData[index:index+length]
    index += length
    return string(ret)
}


func getService() string {
    service :=  strconv.Itoa(int(header.Process)) + strconv.Itoa(int(header.Service)) 
    return service
}
