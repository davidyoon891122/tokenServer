package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"./bodyStruct"
	"./db"
	"./tools"

	"go.mongodb.org/mongo-driver/bson"
)

//Global Variables
var IP string = "10.131.150.171" //docker "172.17.0.2"
var PORT string = ":13302"
var logger *log.Logger
var programName string = "server"

func setLogger() {
	currentDirectory, _ := os.Getwd()
	logPath := "/logs/" + programName + ".log"
	f, err := os.OpenFile(currentDirectory+logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger = log.New(f, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	setLogger()
	ln, err := net.Listen("tcp", IP+PORT)
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
	var parsedData *bodyStruct.Body
	client := conn.RemoteAddr().String()

	logger.Printf("client %s is connected...", client)
	defer fmt.Printf("client %s is disconnected...\n", client)
	for {
		recvBuf = make([]byte, 4096)
		n, err := conn.Read(recvBuf)

		if err != nil {
			if err == io.EOF {
				continue
			} else {
				logger.Printf("conn.Read error; error : %v", err)
			}
		}

		if n > 0 {
			data := recvBuf[:n]
			parsedData = tools.Parse(data)
			logger.Printf("data from client : %v", string(data))

		}

		resCode := db.WriteData(parsedData)
		if resCode == 0 {
			fmt.Println("UserInfo succesfully saved ")
			savedData := db.ReadData("")

			for k, v := range savedData.([]bson.M) {
				fmt.Printf("%d. UserID : %s\n", k, v["userID"])
				fmt.Printf("%d. Token : %s\n", k, v["token"])
				fmt.Printf("-------------------------------------\n")
			}

		} else if resCode == 11000 {
			fmt.Println("Existed ID")
			result := db.ReadData(parsedData.UserID)
			fmt.Println("Existed data : ", result)
		}

	}
}
