package openwtester

import (
	"github.com/blocktree/openwallet/common/file"
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
	"github.com/blocktree/openwallet/openwallet"
)

var (
	testApp        = "assets-adapter"
	configFilePath = filepath.Join("conf")
	dbFilePath = filepath.Join("data", "db")
	dbFileName = "blockchain-eos.db"

	tw *openw.WalletManager
)

func init() {
	tw = testInitWalletManager()
}


func testInitWalletManager() *openw.WalletManager {
	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = configFilePath
	tc.EnableBlockScan = false
	tc.SupportAssets = []string{
		"EOS",
	}

	file.MkdirAll(dbFilePath)

	return openw.NewWalletManager(tc)
	//tm.Init()
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := tw
	w := &openwallet.Wallet{Alias: "HELLO EOS", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("wallet:", nw)
	log.Info("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := tw

	wallet, err := tm.GetWalletInfo(testApp, "W7tue6SDce38fPwerdKqyebUh6yo2nTQLC")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	log.Info("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := tw

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("wallet[", i, "] :", w)
	}
	log.Info("wallet count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := tw

	walletID := "WEyoXkvytkkbK7RJLdoS4H7hbdjDAvRXjY"
	account := &openwallet.AssetsAccount{Alias: "hrt3arlcl354", WalletID: walletID, Required: 1, Symbol: "EOS", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("account:", account)
	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := tw

	walletID := "WEyoXkvytkkbK7RJLdoS4H7hbdjDAvRXjY"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("account[", i, "] :", w)
	}
	log.Info("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := tw

	walletID := "WGVsUfTTVaCwAMRTqeJiDQsZ3vrWp9DzMA"
	accountID := "CbnmpvJNsUjtEMRoy5Nf5FGTyfjLbke8FuKjKtEUc7fs"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 1)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "VzMPgLhU83HbsGabLDAzyAM5NaY14d6gZc"
	accountID := "2UdnebpnSAk8b5btBD5e5dAQLcankywAX3ACaz7deim3"
	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("address[", i, "] :", w.Address)
	}
	log.Info("address count:", len(list))

	tm.CloseDB(testApp)
}
