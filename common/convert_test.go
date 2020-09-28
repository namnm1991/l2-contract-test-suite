package common

import (
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUint48ToBytes(t *testing.T) {
	var x uint64 = 1
	data := hexutil.Encode(Uint48ToBytes(x))
	require.Equal(t, "0x000000000001", data)
}

func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		pubKey := make([]byte, 32)
		rand.Read(pubKey)
		fmt.Println(hexutil.Encode(pubKey))
	}
}
