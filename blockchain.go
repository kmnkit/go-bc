package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Block 블록 기본 구조체
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// NewBlock 새 블록 생성
func NewBlock(nonce int, previousHash string) *Block {

	// Block의 address를 반환함
	// 방법 1
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b

	// 방법 2
	// return &Block{
	// 	timestamp: time.Now().UnixNano(), // UnixNano는 int64타입의 값을 반환함.
	// }
}

// Print 블록의 정보를 출력하는 메소드
func (b *Block) Print() {
	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previousHash   %s\n", b.previousHash)
	fmt.Printf("transactions   %s\n", b.transactions)
}

// Blockchain 블록체인 구조체
type Blockchain struct {
	transactionPool []string
	chain           []*Block // 이 체인에 블록을 추가해나감.
}

// NewBlockChain 새 블록체인 작성
func NewBlockChain() *Blockchain {
	/*
		새 블록체인을 작성한다.
		처음 블록에는 PreviousHash가 없으므로 InitHash를 작성하여 넣어준다.
		그 블록체인을 넘겨준다.
	*/
	bc := new(Blockchain)
	bc.CreateBlock(0, "InitHash")
	return bc
}

// CreateBlock 블록을 하나 추가하여 블록체인에 그 블록을 추가함
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash) // 블록 작성
	bc.chain = append(bc.chain, b)     // bc의 체인에 위의 블록을 추가
	return b
}

// Print 블록 체인 안의 체인 안의 모든 블록을 출력함
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 60))
	return
}

func init() {
	log.SetPrefix("BlockChain: ")
}

func main() {
	blockChain := NewBlockChain()
	blockChain.Print()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
	return
}
