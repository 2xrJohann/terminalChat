package main

import (
    "fmt"
    "net"
    "bufio"
    "sync"
    "strconv"
)

type Client struct {
    reader     *bufio.Reader
    writer     *bufio.Writer
    conn       net.Conn
    name	   string
}

func broadcast(connections *[]Client, text *string, sender string){
	var temp = *connections
	for i2 := 0; i2 < len(temp); i2++{
		if temp[i2].name != sender {
		fmt.Fprintf(temp[i2].conn, "Anon:" + (*text))
		}
	}
}

func connect(conn net.Conn, connections *[]Client, wg sync.WaitGroup){
	wg.Add(1)
	tempReader := bufio.NewReader(conn)
	tempWriter := bufio.NewWriter(conn)
	fmt.Fprintf(conn, "Welcome!\n")
	randy := strconv.Itoa((len(*connections)))

	tempClient := Client{reader: tempReader, writer: tempWriter, conn: conn, name: randy}

	temp := *connections
	temp = append(temp, tempClient)
	*connections = temp

	fmt.Println("--New connection!--")
	for{
		listenMsg(*tempReader, connections, randy)
	}
}

func listen(wg sync.WaitGroup, connections *[]Client){
	ln, _ := net.Listen("tcp", ":8000")
	defer wg.Done()

  	for {
  		conn, _ := ln.Accept()
  		go connect(conn, connections, wg)
  	}
}

func listenMsg(stream bufio.Reader, connections *[]Client, name string){
	for{
		message, _ := stream.ReadString('\n')
		msgpointer := &message
    	broadcast(connections, msgpointer, name)
    }
}

func main() {
	var connections []Client

    message := "--Start Server--"
    fmt.Println(message)

    var wg sync.WaitGroup
    wg.Add(1)

    go listen(wg, &connections)
    wg.Wait()
}

