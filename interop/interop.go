package interop

import (
	"github.com/btcsuite/btcd/wire"
	"github.com/ethereum/go-ethereum/core/types"
)

func ConvBlock(block *wire.MsgBlock) (*types.Block, error) {

	return nil, nil
}

func ConvBlockHeader(header *wire.BlockHeader) (*types.Header, error) {

	return nil, nil
}

// Transaction types.
const (
	LegacyTxType     = types.LegacyTxType
	AccessListTxType = types.AccessListTxType
	DynamicFeeTxType = types.DynamicFeeTxType
	BlobTxType       = types.BlobTxType
	BTCTxType        = 0x04
)

func ConvReceipts(tx *wire.MsgTx) (*types.Receipt, error) {

	return nil, nil
}
