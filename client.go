package main

import(
  "net"
  "fmt"
  "bufio"
  "os"
)


func listenMsg(stream bufio.Reader){
  for{
    message, _ := stream.ReadString('\n')
      fmt.Print(string(message)) 
    }
}

func send(conn net.Conn, reader bufio.Reader){
  text, _ := reader.ReadString('\n')
  fmt.Fprintf(conn, text)
}

func main() {
  conn, _ := net.Dial("tcp", "127.0.0.1:8000")
  reader := bufio.NewReader(os.Stdin)
  receiver := bufio.NewReader(conn)
  go listenMsg(*receiver)
  for {
    send(conn, *reader)
  }

}