package types

import (
	"crypto/rand"
	"fmt"
)

type Hash [32]uint8

func HashFromBytes(bytes []byte) Hash {
	if len(bytes) != 32 { 
		panic(fmt.Sprintf("the provided bytes len is not equal to 32"))
	}

	var value [32]uint8

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