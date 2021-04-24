package main

import (
	"fmt"
	"log"
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

func init() {
	log.SetPrefix("BlockChain: ")
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
