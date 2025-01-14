/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract, extParam map[string]interface{}) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "", contract, extParam)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract,
	feeSupportAccount *openwallet.FeesSupportAccount) ([]*openwallet.RawTransactionWithError, error) {

	rawTxArray, err := tm.CreateSummaryRawTransactionWithError(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract, feeSupportAccount)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer_KLAY(t *testing.T) {

	addrs := []string{
		// "0x513262dc4559347ace1b211bdc85469bb7454744",
		"0x8144a6f2d4873d02256640678b402ea233648f15",
		//"0x48740446f5637995b3b542832ba8a511caeafaa4",
		//"0x9fad88195e6ee7f8c39e9e4ed4deb70a21836ada",
		//"0xa02126f69d4e240ef4e373224b11f0dbaf652c76",
		//"0xf1dd51bdb6234b8d9154bb73f55ac9683166a733",
		//"0xf41fbb39d2d57de11b065dffe4d9c5fb535e25ed",

		//"0x0220655ae9f32a291d00e7cf1cecc9f2b7964f00",
	}

	tm := testInitWalletManager()
	walletID := "W8tvfXwczfjJmwSoVwQkKqzDnhB72GzP9Y"
	accountID := "4a8rThPaKF7ZCobSicZQ6abP1c53dgZPiq1f3CeSt4TJ"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	for _, to := range addrs {
		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "1.99", "", nil, nil)
		if err != nil {
			return
		}

		log.Std.Info("rawTx: %+v", rawTx)

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

	}
}

func TestTransfer_ERC20(t *testing.T) {

	addrs := []string{
		"0x8144a6f2d4873d02256640678b402ea233648f15",
		// "0x3880f535ea2e5ea837d4f72250ede40627ccdca0",
		// "0x48740446f5637995b3b542832ba8a511caeafaa4",
		// "0x9fad88195e6ee7f8c39e9e4ed4deb70a21836ada",
		// "0xa02126f69d4e240ef4e373224b11f0dbaf652c76",
		// "0xf1dd51bdb6234b8d9154bb73f55ac9683166a733",
		// "0xf41fbb39d2d57de11b065dffe4d9c5fb535e25ed",
	}

	tm := testInitWalletManager()
	walletID := "W8tvfXwczfjJmwSoVwQkKqzDnhB72GzP9Y"
	accountID := "4a8rThPaKF7ZCobSicZQ6abP1c53dgZPiq1f3CeSt4TJ"

	contract := openwallet.SmartContract{
		Address:  "0x437514dbb392af27b5daaf98030ab6f028a49513",
		Symbol:   "KLAY",
		Name:     "KERRI",
		Token:    "KERRI",
		Decimals: 2,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	for _, to := range addrs {
		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "1.23", "", &contract, nil)
		if err != nil {
			return
		}

		log.Std.Info("rawTx: %+v", rawTx)

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

	}

}

func TestSummary_KLAY(t *testing.T) {
	tm := testInitWalletManager()

	walletID := "W8tvfXwczfjJmwSoVwQkKqzDnhB72GzP9Y"
	accountID := "4a8rThPaKF7ZCobSicZQ6abP1c53dgZPiq1f3CeSt4TJ"
	summaryAddress := "0xf5f6c07361826def0d7f463816b6c983815063f6"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, nil, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}

func TestSummary_ERC20(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W8tvfXwczfjJmwSoVwQkKqzDnhB72GzP9Y"
	accountID := "4a8rThPaKF7ZCobSicZQ6abP1c53dgZPiq1f3CeSt4TJ"
	summaryAddress := "0xf5f6c07361826def0d7f463816b6c983815063f6"

	// feesSupport := openwallet.FeesSupportAccount{
	// 	AccountID: "4a8rThPaKF7ZCobSicZQ6abP1c53dgZPiq1f3CeSt4TJ",
	// 	//FixSupportAmount: "0.01",
	// 	FeesSupportScale: "1.3",
	// }

	contract := openwallet.SmartContract{
		Address:  "0x437514dbb392af27b5daaf98030ab6f028a49513",
		Symbol:   "KLAY",
		Name:     "KERRI",
		Token:    "KERRI",
		Decimals: 2,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, &contract, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}
