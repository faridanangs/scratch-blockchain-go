package types

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

func GenerateAddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		panic(fmt.Sprintf("given bytes with length %d should be 20", len(b)))
	}

	var value [20]uint8
	for i := range 20 {
		value[i] = b[i]
	}

	fmt.Println("Address after: ", value)
	fmt.Println("Address after convert: ", Address(value))

	return Address(value)

}

func (a *Address) ToSlice() []byte {
	b := make([]byte, 20)

	for i := range 20 {
		b[i] = a[i]
	}

	return b
}

func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}
