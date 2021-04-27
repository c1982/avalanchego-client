package api

type MintNFT struct {
	Assetid    string   `json:"assetID"`
	Payload    string   `json:"payload"`
	To         string   `json:"to"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type NFTAsset struct {
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	Mintersets []struct {
		Minters   []string `json:"minters"`
		Threshold int      `json:"threshold"`
	} `json:"minterSets"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

//CreateNFTAsset Create a new non-fungible asset. No units of the asset exist at initialization. Minters can mint units of this asset using avm.mintNFT
func (c Calls) CreateNFTAsset(asset NFTAsset) (assetID, changeAddr string, err error) {
	rsp, err := c.client.NewRequestStruct("ext/bc/X", "amv.createVariableCapAsset", asset)
	if err != nil {
		return assetID, changeAddr, err
	}
	err = rsp.
		OutStr("assetID", &assetID).
		OutStr("changeAddr", &changeAddr).
		Error()

	return assetID, changeAddr, err
}

//MintNFT Mint non-fungible tokens which were created with avm.createNFTAsset.
func (c Calls) MintNFT(asset MintNFT) (txID, changeAddr string, err error) {
	rsp, err := c.client.NewRequestStruct("ext/bc/X", "avm.mintNFT", asset)
	if err != nil {
		return txID, changeAddr, err
	}
	err = rsp.
		OutStr("txID", &txID).
		OutStr("changeAddr", &changeAddr).
		Error()

	return txID, changeAddr, err
}
