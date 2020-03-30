package Peer

import "fmt"
import "net"
import "log"
import "sync"
import a1 "github.com/shahrukhx04/assignment01IBC"
// import "encoding/gob"


func HandleConnection(c net.Conn, BlockChain *a1.Block){

  fmt.Println("-----------------------------------------------------------------------")
  fmt.Println("Server Responded", c.RemoteAddr())

  recvdSlice := make([]byte, 100)
  c.Read(recvdSlice)
  fmt.Println("Data Recieved from the server")
  fmt.Println(string(recvdSlice))

  // dec := gob.NewDecoder(c)
  // var DecodedData = dec.Decode(&BlockChain)
  // _ = DecodedData
  fmt.Println("-----------------------------------------------------------------------")
}

func MinorValidation(c net.Conn, PortNumber string, Chain *a1.Block){

    fmt.Println("-----------------------------------------------------------------------")
    fmt.Println(PortNumber + " Transaction Recieved from Node", c.RemoteAddr())
    recvdSlice := make([]byte, 100)
    c.Read(recvdSlice)
    fmt.Println(string(recvdSlice) + "Transaction")

    fmt.Println("I am validating the Block")

    fmt.Println("Block Validated Successfully")

    fmt.Println("I am mining the block")

    var CurrentTrans a1.MakeTransaction

    CurrentTrans.TransactionID = 1
    CurrentTrans.SenderToRecieverAmount = 0
    CurrentTrans.SenderName = "GensisBlock"
    CurrentTrans.RecieverName = "GensisBlock"


    a1.InsertBlock(string(recvdSlice),Chain, CurrentTrans)
    fmt.Println("Block Mined Successfully")
    fmt.Println("I am sending mined block o all other nodes for validation")
    SendBlockToOtherNodesForValidation(PortNumber,Chain)

    fmt.Println("-----------------------------------------------------------------------")


}

func Listen(PortNumber string, wg *sync.WaitGroup, chainHead *a1.Block){

    conn, err := net.Dial("tcp", "localhost:9999")

    var BlockChain *a1.Block

    if err != nil {
    // handle error

    }
    recvdSlice := make([]byte, 100)
    conn.Read(recvdSlice)
    fmt.Println("Data Recieved from the server")
    fmt.Println(string(recvdSlice))


    In, err := net.Listen("tcp", ":"+PortNumber)
    if err != nil {
      log.Fatal(err)
    }



    for i:=0; i<5 ; i++{

        conn1, err1 := In.Accept()
        if err1 != nil {
          fmt.Println(err1)
        }
        HandleConnection(conn1, BlockChain)
    }



    for i:=0; ; i++{

        fmt.Println( PortNumber + " Now accepting the transactions")
        conn1, err1 := In.Accept()
        if err1 != nil {
          fmt.Println(err1)
        }
        MinorValidation(conn1, PortNumber, chainHead)
    }
    defer wg.Done()

}

func SendBlockToOtherNodesForValidation(SenderPort string, BlockChain *a1.Block){

  ArrayOfPortNumbers:= [5]string{"1234","2345","3456","4567","5678"}
  for i:=0; i < 5; i++{
    if ArrayOfPortNumbers[i] != SenderPort {
      // conn1, err1 := net.Dial("tcp", "localhost:" + ArrayOfPortNumbers[i])
      // conn1.Write([]byte(SenderPort + " am Sending Block to  " + ArrayOfPortNumbers[i]))
      // if err1 != nil {
      //   fmt.Println(err1)
      // }
    }
  }


}


func ConnectWithOtherClients(PortNumber string){

  fmt.Println("I am dailing to other nodes")
  ArrayOfPortNumbers:= [5]string{"1234","2345","3456","4567","5678"}
  for i:=0; i < 5; i++{
    if ArrayOfPortNumbers[i] != PortNumber {
      conn1, err1 := net.Dial("tcp", "localhost:" + ArrayOfPortNumbers[i])
      conn1.Write([]byte(PortNumber + " am connecting with " + ArrayOfPortNumbers[i]))
      if err1 != nil {
        fmt.Println(err1)
      }
    }
  }
}
