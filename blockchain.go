package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

// Block 블록 기본 구조체
type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}

// NewBlock 새 블록 생성
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	// Block의 address를 반환함
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

// MarshalJSON JSON 형태 지정
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

// Print 블록의 정보를 출력하는 메소드
func (b *Block) Print() {
	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previousHash   %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

// Blockchain 블록체인 구조체
type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block // 이 체인에 블록을 추가해나감.
	blockchainAddress string
}

// NewBlockChain 새 블록체인 작성
func NewBlockChain(blockchainAddress string) *Blockchain {
	// 새 블록체인을 작성한여 넘겨준다.
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

// CreateBlock 블록을 하나 추가하여 블록체인에 그 블록을 추가함
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool) // 블록 작성
	bc.chain = append(bc.chain, b)                         // bc의 체인에 위의 블록을 추가
	bc.transactionPool = []*Transaction{}
	return b
}

// LastBlock 블록체인의 마지막 블록을 구함
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
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

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

// Transaction 트랜잭셩 구조체
type Transaction struct {
	senderBlockchainAddress    string  // 보내는이 Address
	recipientBlockchainAddress string  // 받는이 Address
	value                      float32 // 보내는 값
}

// AddTransaction 블록의 transactionPool에 새 트랜잭션을 추가함
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
	return
}

// CopyTransactionPool 트랜잭션 풀을 복사해둠
func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(
			t.senderBlockchainAddress,
			t.recipientBlockchainAddress,
			t.value,
		))
	}
	return transactions
}

// ValidProof 계산 결과 검증
func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

// ProofOfWork 해를 구함
func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

// NewTransaction 새 트랜잭션 작성
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address       %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address    %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                           %.1f\n", t.value)
}

// MarshalJSON Transaction 내용을 Json Marshal 함.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

func init() {
	log.SetPrefix("BlockChain: ")
}

func main() {
	myBlockchainAddress := "my_blockchain_address"

	blockChain := NewBlockChain(myBlockchainAddress)
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	blockChain.Mining()
	blockChain.Print()
	return
}
