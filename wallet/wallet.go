package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

// NewWallet 새 월렛 작성
func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey
	return w
}

// PrivateKey
func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

// PrivateKeyStr
func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

// PublicKey
func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

// PublicKeyStr
func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}
