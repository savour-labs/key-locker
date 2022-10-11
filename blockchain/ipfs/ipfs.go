package ipfs

import (
	"github.com/savour-labs/key-locker/blockchain"
	"github.com/savour-labs/key-locker/blockchain/fallback"
	"github.com/savour-labs/key-locker/blockchain/multiclient"
	"github.com/savour-labs/key-locker/config"
	"github.com/savour-labs/key-locker/proto/common"
	"github.com/savour-labs/key-locker/proto/keylocker"
)

const ChainName = "Ipfs"

type KeyAdaptor struct {
	fallback.KeyAdaptor
	clients *multiclient.MultiClient
}

func NewChainAdaptor(conf *config.Config) (blockchain.KeyAdaptor, error) {
	return &KeyAdaptor{
		clients: nil,
	}, nil
}

func NewLocalKeyAdaptor(network config.NetWorkType) blockchain.KeyAdaptor {
	return newKeyAdaptor()
}

func newKeyAdaptor() blockchain.KeyAdaptor {
	return &KeyAdaptor{
		clients: multiclient.New([]multiclient.Client{}),
	}
}

func (a *KeyAdaptor) getClient() {

}

func (a *KeyAdaptor) GetSupportChain(req *keylocker.SupportChainReq) (*keylocker.SupportChainRep, error) {
	return &keylocker.SupportChainRep{
		Code: common.ReturnCode_SUCCESS,
		Msg:  "get support chain success",
	}, nil
}

func (a *KeyAdaptor) GetSocialKey(req *keylocker.GetSocialKeyReq) (*keylocker.GetSocialKeyRep, error) {
	// todo: 调用 IPFS 接口获取 Social Key
	return &keylocker.GetSocialKeyRep{
		Code: common.ReturnCode_SUCCESS,
		Msg:  "get ipfs social key success",
	}, nil
}

func (a *KeyAdaptor) SetSocialKey(req *keylocker.SetSocialKeyReq) (*keylocker.SetSocialKeyRep, error) {
	// todo: 调用 IPFS 接口上传 Social Key
	return &keylocker.SetSocialKeyRep{
		Code: common.ReturnCode_SUCCESS,
		Msg:  "set ipfs social key success",
	}, nil
}
