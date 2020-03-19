package main

import (
    "fmt"
	"log"
	"net"
	"os"
    "./tools"
    "./db"
    "./bodyStruct"
)

//Global Variables
var IP string = "172.17.0.2"
var PORT string = ":13302"
var logger *log.Logger
var programName string = "Server"

func setLogger() {
    currentDirectory, _ := os.Getwd()
    logPath := "/logs/" + programName + ".log"
	f, err := os.OpenFile(currentDirectory + logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
        panic(err)
	}

	logger = log.New(f, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
}



func main() {
	setLogger()
	ln, err := net.Listen("tcp", IP + PORT)
	defer ln.Close()

	if err != nil {
		logger.Fatalf("net.Listen error; error : %v", err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			logger.Fatalf("ln.Accept error; error : %v", err)
			continue
		}

		go Handler(conn)
	}
}


func Handler(conn net.Conn) {
	var recvBuf []byte
    var parsedData bodyStruct.Body
	client := conn.RemoteAddr().String()

	logger.Printf("client %s is connected...", client)

	recvBuf = make([]byte, 4096)
	n, err := conn.Read(recvBuf)

	if err != nil {
		logger.Printf("conn.Read error; error : %v", err)
	}

	if n > 0 {
		data := recvBuf[:n]
        parsedData = tools.Parse(data)
        logger.Printf("data from client : %v", string(data))

	}

    savedData := db.ReadData()

    fmt.Println("saved data :", savedData)

    db.WriteData(parsedData) 
    defer fmt.Printf("client %s is disconnected...\n", client)
}






