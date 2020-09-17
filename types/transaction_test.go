package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFee_MarshalText(t *testing.T) {
	fee := &Fee{
		Mantisa: 34,
		Exp:     1,
	}
	b, err := fee.MarshalText()
	require.NoError(t, err)
	require.Equal(t, string(b), "0x154")
}

func TestSettlementOp1 (t *testing.T){
	op := Settlement1{
		opType:       SettlementOp11,
		Token1:       1,
		Token2:       2,
		Account1:     14,
		Account2:     15,
		Rate1:        Amount{
			Mantisa: 1,
			Exp:     18,
		},
		Rate2:        Amount{
			Mantisa: 1,
			Exp:     18,
		},
		Amount1:      Amount{
			Mantisa: 2,
			Exp: 14,
		},
		Amount2:      Amount{
			Mantisa: 3,
			Exp: 14,
		},
		Fee1:         Fee{
			Mantisa: 1,
			Exp:     6,
		},
		Fee2:         Fee{
			Mantisa: 1,
			Exp: 6,
		},
		ValidSince1:  1600331441,
		ValidSince2:  1600331441,
		ValidPeriod1: 86400,
		ValidPeriod2: 86400,
	}
	b, err :=json.MarshalIndent(&op, "", "")
	require.NoError(t, err)
	t.Log(string(b))
}
