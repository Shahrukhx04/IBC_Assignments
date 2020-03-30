package main
import a1 "github.com/shahrukhx04/assignment01IBC"
// import a2 "github.com/shahrukhx04/assignment01IBC/Assignment_2_Peer"
import "fmt"
import "net"
import "log"
import "sync"
import "math/rand"
import "time"
// import "encoding/gob"


func HandleConnection( c net.Conn){

  fmt.Println("-----------------------------------------------------------------------")
  fmt.Println("Client has conencted", c.RemoteAddr())
  c.Write([]byte("You want to become rich, wait"))
  fmt.Println("-----------------------------------------------------------------------")
}

func SentDataToClient(PortNumber string, chainHead *a1.Block,  wg *sync.WaitGroup, c net.Conn){
  conn1, err1 := net.Dial("tcp", "localhost:" + PortNumber)
  conn1.Write([]byte("Server Responding to Client: " + PortNumber))
  if err1 != nil {
    fmt.Println(err1)
  }

  // enc := gob.NewEncoder(conn1)
  // var EncodedData = enc.Encode(chainHead)
  // _ = EncodedData
  defer wg.Done()
}


func MakeATransaction(SenderPort string, wg *sync.WaitGroup){
  var Trasaction string = "Send50CoinsToPort:" + SenderPort

  var CurrentTrans a1.MakeTransaction
  CurrentTrans.TransactionID = 10
  CurrentTrans.SenderToRecieverAmount = 10
  CurrentTrans.SenderName = "Satoshi"
  CurrentTrans.RecieverName = "Alice"

  var Minor = rand.Intn(4 - 0) + 0
  var ArrayOfMiners = [5]string{"2345","3456","4567","5678"}


  conn1, err1 := net.Dial("tcp", "localhost:" + ArrayOfMiners[Minor])
  conn1.Write([]byte(Trasaction))
  fmt.Println("Transaction sent to " + SenderPort)
  if err1 != nil {
    fmt.Println(err1)
  }
  defer wg.Done()
}

func main(){

  var conn net.Conn

  In, err := net.Listen("tcp", ":9999")


  var AddressOfClients [5]string;

  var chainHead *a1.Block

  var CurrentTrans a1.MakeTransaction
  CurrentTrans.TransactionID = 3
  CurrentTrans.SenderToRecieverAmount = 100
  CurrentTrans.SenderName = "Satoshi"
  CurrentTrans.RecieverName = "Satoshi"


  chainHead = a1.InsertBlock("GenesisBlock: Pay 100 coins to Satoshi", nil, CurrentTrans)
  if err != nil {
    log.Fatal(err)
  }

  var wg sync.WaitGroup
  wg.Add(5)

  for i:=0; i<5 ; i++{
    fmt.Println("Satoshi is Listening")
    conn, err := In.Accept()

    CurrentTrans.TransactionID = 4+i
    CurrentTrans.SenderToRecieverAmount = 100
    CurrentTrans.SenderName = conn.RemoteAddr().String()
    CurrentTrans.RecieverName = "Satoshi"
    chainHead = a1.InsertBlock("Pay 100 coins to Satoshi from " + conn.RemoteAddr().String(), chainHead, CurrentTrans)
    AddressOfClients[i] = conn.RemoteAddr().String()
    if err != nil {
      fmt.Println(err)
    }
    HandleConnection(conn)
  }

  a1.ListBlocks(chainHead)

  go SentDataToClient("1234", chainHead, &wg, conn)
  go SentDataToClient("2345", chainHead, &wg, conn)
  go SentDataToClient("3456", chainHead, &wg, conn)
  go SentDataToClient("4567", chainHead, &wg, conn)
  go SentDataToClient("5678", chainHead, &wg, conn)

  wg.Wait()


  time.Sleep(10*time.Second)
  var wg1 sync.WaitGroup
  wg1.Add(1)

  go MakeATransaction("1234",&wg1)

  wg1.Wait()

}
