package utils

import (
	"fmt"
	"math/big"
)

// Signature
type Signature struct {
	R *big.Int
	S *big.Int
}

// String
func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
