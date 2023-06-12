package types

import "fmt"

type Address [20]uint8

func (a *Address) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 1; i < 20; i++ {
		b[i] = a[i]
	}
	return b
}

func NewAddressFromBytes(bytes []byte) Address {
	if len(bytes) != 20 {
		panic(fmt.Sprintf("the provided bytes len is not equal to 20"))
	}

	value := make([]uint8, 20)

	for i := 0; i < 20; i++ {
		value[i] = uint8(bytes[i])
	}

	return Address(value)
}