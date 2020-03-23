package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	var IP string = "10.131.150.171"
	var PORT string = ":13302"

	conn, err := net.Dial("tcp", IP+PORT)

	defer conn.Close()
	if err != nil {
		panic(err)
	}

	for {
		printMenu()
		data := selectMenu()
		if data == nil {
			continue
		}
		conn.Write(data)
	}
}

func packMsg(length int, process int, service int, window int, control int, flag int, reserve int) []byte {
	var headerBuffer bytes.Buffer
	var headerBytes []byte
	var bodyBytes []byte
	var totalBytes []byte
	//header
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(length))
	headerBuffer.Write(buf)
	buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(process))
	headerBuffer.Write(buf)
	buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(service))
	headerBuffer.Write(buf)
	buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(window))
	headerBuffer.Write(buf)
	headerBuffer.WriteByte(byte(control))
	headerBuffer.WriteByte(byte(flag))
	buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(reserve))
	headerBuffer.Write(buf)

	headerBytes = headerBuffer.Bytes()
	//body
	if service == 0x0000 {
		bodyBytes = packTestBody()
	} else if service == 0x0001 {
		bodyBytes = packInsertBody()
	} else if service == 0x0002 {
		bodyBytes = packInsertBody()
	}

	totalBytes = append(headerBytes, bodyBytes...)
	return totalBytes
}

func packTestBody() []byte {
	var token string
	var userID string
	var bodyBuffer bytes.Buffer

	token = "eR94gyvABa8:APA91bHpDrn3xN3CqWPNfybP3ztK-_QJUi0lksxXcPcaaoCyKCD96CIP3a7SzlYb68a2yy9a2TLCK8SVFOcNvSGkceN3PJBZeXQXF15oekzI_whVMmi5Sk-Xtn6tZ3XrvlebhbF1oTlb"
	userID = "htsadmin"
	bodyBuffer.Write([]byte(token))
	bodyBuffer.Write([]byte(userID))

	return bodyBuffer.Bytes()
}

func packInsertBody() []byte {
	var token string
	var userID string
	var bodyBuffer bytes.Buffer

	token = getToken()
	fmt.Println("User ID : ")
	fmt.Scanf("%s", &userID)
	bodyBuffer.Write([]byte(token))
	bodyBuffer.Write([]byte(userID))

	return bodyBuffer.Bytes()
}

func printMenu() {
	fmt.Println("1. Duplicated key test.")
	fmt.Println("2. Insert Test.")
	fmt.Println("3. Exit.")
}

func selectMenu() []byte {
	var selector int
	fmt.Scanf("%d", &selector)

	//data packer
	var length int
	var process int
	var service int
	var window int
	var control int
	var flag int
	var reserve int
	var data []byte

	if selector == 1 {
		length = 152 + 16
		process = 0x7FF3
		service = 0x0000
		window = 1050
		control = 0
		flag = 0
		reserve = 0
		data = packMsg(length, process, service, window, control, flag, reserve)
	} else if selector == 2 {
		length = 152 + 16
		process = 0x7FF3
		service = 0x0001
		window = 1050
		control = 0
		flag = 0
		reserve = 0
		data = packMsg(length, process, service, window, control, flag, reserve)
	} else if selector == 3 {
		os.Exit(1)
	} else {
		data = nil
	}

	fmt.Println("data to send :", data)
	return data
}

// func getLength(data string) {
// 	return len(data)
// }

func getToken() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 152)

	now := time.Now().Second()
	rand.Seed(int64(now))

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
