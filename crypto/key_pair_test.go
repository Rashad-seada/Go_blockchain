package crypto

import (
	"fmt"
	"testing"
)

func TestGeneratePairOfKeys(t *testing.T) {
	keypair := GenerateUniqueKeypair()

	PrivateKey := keypair.PrivateKey
	PublicKey := keypair.PublicKey
	Address := keypair.Address()

	fmt.Println(PrivateKey)

	fmt.Println(PublicKey)

	fmt.Println(Address)
}