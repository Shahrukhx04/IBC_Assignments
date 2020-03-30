package Assignment_1


import (
    "fmt"
    "encoding/json"
    "crypto/sha256"
)

type MakeTransaction struct{
  SenderName string
  RecieverName string
  TransactionID int
  SenderToRecieverAmount int
}

type Block struct {
    Trasaction string
    Previous *Block
    HashData [32]byte
    BlockChainTransaction MakeTransaction
}


func InsertBlock(Transaction string, chainHead *Block, BlockChainTrans MakeTransaction) *Block {
      var InsertedBlock Block
      InsertedBlock.Trasaction = Transaction
      InsertedBlock.Previous = chainHead

      InsertedBlock.BlockChainTransaction.RecieverName = BlockChainTrans.RecieverName
      InsertedBlock.BlockChainTransaction.SenderName = BlockChainTrans.SenderName
      InsertedBlock.BlockChainTransaction.SenderToRecieverAmount = BlockChainTrans.SenderToRecieverAmount
      InsertedBlock.BlockChainTransaction.TransactionID = BlockChainTrans.TransactionID

      out, err := json.Marshal(InsertedBlock.Previous)
      _ = err
      var Headers[] byte= []byte(InsertedBlock.Trasaction + string(out))
      InsertedBlock.HashData = sha256.Sum256([]byte(Headers))
      return &InsertedBlock
}


func ListBlocks(chainHead *Block) {

  for chainHead.Previous!=nil {
    fmt.Println(chainHead.Trasaction)
    chainHead = chainHead.Previous
  }

    fmt.Println(chainHead.Trasaction)
}


func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
  for chainHead.Previous!=nil {
    if chainHead.Trasaction==oldTrans {
      chainHead.Trasaction = newTrans
      return
    }
    chainHead = chainHead.Previous
  }
}

func VerifyChain(chainHead *Block) {

  var InsertedBlock *Block = chainHead

  for InsertedBlock.Previous!=nil {
      fmt.Println("Current Transaction" + InsertedBlock.Trasaction)
      out, err := json.Marshal(InsertedBlock.Previous)
      _ = err
      var Headers[] byte= []byte(InsertedBlock.Trasaction + string(out))
      var CalculatedHastData = sha256.Sum256([]byte(Headers))

      if (CalculatedHastData != InsertedBlock.HashData){
        fmt.Println("Chain is Invalid")
        return;
      }
      InsertedBlock = InsertedBlock.Previous
    }
      fmt.Println("Chain Shahrukh ")

}
