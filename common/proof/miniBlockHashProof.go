package proof

import (
	"fmt"
	util "github.com/KyberNetwork/l2-contract-test-suite/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/KyberNetwork/l2-contract-test-suite/types"
)

func BuildMiniBlockProof(miniBlocks []*types.MiniBlock, miniBlockIndex uint, timeStamp uint32) hexutil.Bytes {
	var miniBlockHashes []common.Hash
	for _, blk := range miniBlocks {
		miniBlockHashes = append(miniBlockHashes, blk.Hash())
	}
	proof := BuildBlockInfoProof(miniBlockHashes, miniBlockIndex)
	proof = append(proof, util.Uint32ToBytes(timeStamp)...)
	proof = append(proof, util.Uint8ToByte(uint8(len(miniBlocks))))
	proof = append(proof, miniBlocks[len(miniBlocks)-1].StateHash.Bytes()...)
	return proof
}

func BuildBlockInfoProof(blks []common.Hash, miniBlockIndex uint) hexutil.Bytes {
	var tmp []common.Hash
	tmp = append(tmp, blks...)

	proof := hexutil.Bytes{0}
	size := len(tmp)
	count := uint8(0)
	for size != 1 {
		for i := 0; i*2 < size; i++ {
			if i*2 == size-1 {
				if uint(i*2) == miniBlockIndex {
					proof = append(proof, common.HexToHash("0x0").Bytes()...)
				}
				tmp[i] = crypto.Keccak256Hash(tmp[i*2].Bytes(), common.HexToHash("0x0").Bytes())
			} else {
				if uint(i*2) == miniBlockIndex {
					proof = append(proof, tmp[i*2+1].Bytes()...)
				}
				if uint(i*2+1) == miniBlockIndex {
					proof = append(proof, tmp[i*2].Bytes()...)
				}
				tmp[i] = crypto.Keccak256Hash(tmp[i*2].Bytes(), tmp[i*2+1].Bytes())
			}
		}
		miniBlockIndex = miniBlockIndex / 2
		size = (size + 1) / 2
		count++
	}
	fmt.Println(tmp[0].Hex())
	proof[0] = util.Uint8ToByte(count)
	return proof
}
