package quorum

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// UnmarshalReceiptJSON unmarshals Receipt from JSON.
func UnmarshalReceiptJSON(input []byte) (r *types.Receipt, err error) {
	type Receipt struct {
		PostState         *hexutil.Bytes  `json:"root"`
		Status            *hexutil.Uint64 `json:"status"`
		CumulativeGasUsed *hexutil.Uint64 `json:"cumulativeGasUsed" gencodec:"required"`
		Bloom             *types.Bloom    `json:"logsBloom"         gencodec:"required"`
		Logs              []*types.Log    `json:"logs"              gencodec:"required"`
		TxHash            *common.Hash    `json:"transactionHash" gencodec:"required"`
		ContractAddress   *common.Address `json:"contractAddress"`
		GasUsed           *hexutil.Uint64 `json:"gasUsed" gencodec:"required"`
		BlockHash         *common.Hash    `json:"blockHash,omitempty"`
		BlockNumber       *hexutil.Big    `json:"blockNumber,omitempty"`
		TransactionIndex  *hexutil.Uint   `json:"transactionIndex"`
	}
	var dec Receipt
	var receipt types.Receipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return nil, err
	}
	if dec.PostState != nil {
		receipt.PostState = *dec.PostState
	}
	if dec.Status != nil {
		receipt.Status = uint64(*dec.Status)
	}
	// if dec.CumulativeGasUsed == nil {
	// 	return nil, errors.New("missing required field 'cumulativeGasUsed' for Receipt")
	// }
	// receipt.CumulativeGasUsed = uint64(*dec.CumulativeGasUsed)
	if dec.Bloom == nil {
		return nil, errors.New("missing required field 'logsBloom' for Receipt")
	}
	receipt.Bloom = *dec.Bloom
	if dec.Logs == nil {
		return nil, errors.New("missing required field 'logs' for Receipt")
	}
	receipt.Logs = dec.Logs
	if dec.TxHash == nil {
		return nil, errors.New("missing required field 'transactionHash' for Receipt")
	}
	receipt.TxHash = *dec.TxHash
	if dec.ContractAddress != nil {
		receipt.ContractAddress = *dec.ContractAddress
	}
	if dec.GasUsed == nil {
		return nil, errors.New("missing required field 'gasUsed' for Receipt")
	}
	receipt.GasUsed = uint64(*dec.GasUsed)
	if dec.BlockHash != nil {
		receipt.BlockHash = *dec.BlockHash
	}
	if dec.BlockNumber != nil {
		receipt.BlockNumber = (*big.Int)(dec.BlockNumber)
	}
	if dec.TransactionIndex != nil {
		receipt.TransactionIndex = uint(*dec.TransactionIndex)
	}
	return &receipt, nil
}
