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
	"path/filepath"
	"testing"

	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/v2/openw"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func TestWalletManager_GetTransactions(t *testing.T) {
	tm := testInitWalletManager()
	list, err := tm.GetTransactions(testApp, 0, -1, "Received", false)
	if err != nil {
		log.Error("GetTransactions failed, unexpected error:", err)
		return
	}
	for i, tx := range list {
		log.Info("trx[", i, "] :", tx)
	}
	log.Info("trx count:", len(list))
}

func TestWalletManager_GetTxUnspent(t *testing.T) {
	tm := testInitWalletManager()
	list, err := tm.GetTxUnspent(testApp, 0, -1, "Received", false)
	if err != nil {
		log.Error("GetTxUnspent failed, unexpected error:", err)
		return
	}
	for i, tx := range list {
		log.Info("Unspent[", i, "] :", tx)
	}
	log.Info("Unspent count:", len(list))
}

func TestWalletManager_GetTxSpent(t *testing.T) {
	tm := testInitWalletManager()
	list, err := tm.GetTxSpent(testApp, 0, -1, "Received", false)
	if err != nil {
		log.Error("GetTxSpent failed, unexpected error:", err)
		return
	}
	for i, tx := range list {
		log.Info("Spent[", i, "] :", tx)
	}
	log.Info("Spent count:", len(list))
}

func TestWalletManager_ExtractUTXO(t *testing.T) {
	tm := testInitWalletManager()
	unspent, err := tm.GetTxUnspent(testApp, 0, -1, "Received", false)
	if err != nil {
		log.Error("GetTxUnspent failed, unexpected error:", err)
		return
	}
	for i, tx := range unspent {

		_, err := tm.GetTxSpent(testApp, 0, -1, "SourceTxID", tx.TxID, "SourceIndex", tx.Index)
		if err == nil {
			continue
		}

		log.Info("ExtractUTXO[", i, "] :", tx)
	}

}

func TestWalletManager_GetTransactionByWxID(t *testing.T) {
	tm := testInitWalletManager()
	wxID := openwallet.GenTransactionWxID(&openwallet.Transaction{
		TxID: "bfa6febb33c8ddde9f7f7b4d93043956cce7e0f4e95da259a78dc9068d178fee",
		Coin: openwallet.Coin{
			Symbol:     "LTC",
			IsContract: false,
			ContractID: "",
		},
	})
	log.Info("wxID:", wxID)
	//"D0+rxcKSqEsFMfGesVzBdf6RloM="
	tx, err := tm.GetTransactionByWxID(testApp, wxID)
	if err != nil {
		log.Error("GetTransactionByTxID failed, unexpected error:", err)
		return
	}
	log.Info("tx:", tx)
}

func TestWalletManager_GetAssetsAccountBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W8tvfXwczfjJmwSoVwQkKqzDnhB72GzP9Y"
	accountID := "4a8rThPaKF7ZCobSicZQ6abP1c53dgZPiq1f3CeSt4TJ"
	//accountID := "3xLbreE3asBRVCCk13Y9V4NzjyijXdx8sb6k54TmQkFg"
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func TestWalletManager_GetAssetsAccountTokenBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W8tvfXwczfjJmwSoVwQkKqzDnhB72GzP9Y"
	//accountID := "CBGfADJdeDDPKh7wywxDrmkTJzxGjAQJyT4hVD44bvLE"
	accountID := "4a8rThPaKF7ZCobSicZQ6abP1c53dgZPiq1f3CeSt4TJ"

	contract := openwallet.SmartContract{
		Address:  "0x437514dbb392af27b5daaf98030ab6f028a49513",
		Symbol:   "KLAY",
		Name:     "KERRI",
		Token:    "KERRI",
		Decimals: 2,
	}

	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance.Balance)
}

func TestWalletManager_GetEstimateFeeRate(t *testing.T) {
	tm := testInitWalletManager()
	coin := openwallet.Coin{
		Symbol: "KLAY",
	}
	feeRate, unit, err := tm.GetEstimateFeeRate(coin)
	if err != nil {
		log.Error("GetEstimateFeeRate failed, unexpected error:", err)
		return
	}
	log.Std.Info("feeRate: %s %s/%s", feeRate, coin.Symbol, unit)
}

func TestGetAddressVerify(t *testing.T) {
	symbol := "KLAY"
	assetsMgr, err := openw.GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}
	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)
	addrDec := assetsMgr.GetAddressDecoderV2()

	flag := addrDec.AddressVerify("0x8144a6f2d4873d02256640678b402ea233648f1")
	log.Infof("flag: %v, expect: false", flag)

	flag = addrDec.AddressVerify("xdfe6070e1e8e53d23147df1e7ed09f49acbbf722")
	log.Infof("flag: %v, expect: false", flag)

	flag = addrDec.AddressVerify("0x513262dc4559347ace1b211bdc85469bb7454744")
	log.Infof("flag: %v, expect: true", flag)

}
