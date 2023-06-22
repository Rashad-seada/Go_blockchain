package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePairOfKeys(t *testing.T) {
	keypair1 := GeneratePrivateKey()

	PrivateKey := keypair1.key
	PublicKey := keypair1.PublicKey()
	Address := keypair1.PublicKey().Address()

	fmt.Println(">>",PrivateKey)
	fmt.Println(">>",PublicKey)
	fmt.Println(">>",Address)
}

func TestKeySignatureVerify(t *testing.T) {
	keypair1 := GeneratePrivateKey()
	keypair2 := GeneratePrivateKey()

	msg := []byte("Hello world")
	sig , err := keypair1.Sign(msg)

	assert.Nil(t,err)
	fmt.Println(">>",sig)

	assert.True(t,sig.Verify(keypair1.PublicKey(),msg,))
	assert.False(t,sig.Verify(keypair2.PublicKey(),msg))
}

