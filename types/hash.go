package types

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Hash [32]uint8

func (h *Hash) ToSlice() []byte {
	b := make([]byte,32)
	for i:= 1; i < 32;i++{
		b[i] = h[i]
	}
	return b
}

func (h *Hash) ToString() string {
	return hex.EncodeToString(h.ToSlice())
}

func (h *Hash) IsZero() bool {
	for i:= 0 ; i < 32 ; i++ {
		if h[i] != 0 {
			return false
		}	
	}
	return true
}
	

func HashFromBytes(bytes []byte) Hash {
	if len(bytes) != 32 { 
		panic(fmt.Sprintf("the provided bytes len is not equal to 32"))
	}

	value := make([]uint8, 32) 

	for i:= 0; i < 32; i++ {
		value[i] = uint8(bytes[i])
	}

	return Hash(value)
}

func RandomByte(size int) []byte {
	token := make([]byte,size)
	rand.Read(token)
	return token
}

func RandomHash(size int) Hash {
	return HashFromBytes(RandomByte(size))
}