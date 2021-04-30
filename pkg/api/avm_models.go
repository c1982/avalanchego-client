package api

//FixedCapAssetArgs CreateFixedCapAsset	signature
type FixedCapAssetArgs struct {
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Denomination   int    `json:"denomination"`
	Initialholders []struct {
		Address string `json:"address"`
		Amount  int    `json:"amount"`
	} `json:"initialHolders"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type AvmMintArgs struct {
	Amount     int64    `json:"amount"`
	Assetid    string   `json:"assetID"`
	To         string   `json:"to"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type CreateVariableCapAssetArgs struct {
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

type ExportArgs struct {
	To         string   `json:"to"`
	Amount     int      `json:"amount"`
	Assetid    string   `json:"assetID"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type ExportAVAXArgs struct {
	To         string   `json:"to"`
	Amount     int      `json:"amount"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type MintNFTArgs struct {
	Assetid    string   `json:"assetID"`
	Payload    string   `json:"payload"`
	To         string   `json:"to"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type NFTAssetArgs struct {
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

type BalancesResponse struct {
	Balances []struct {
		Asset   string `json:"asset"`
		Balance string `json:"balance"`
	} `json:"balances"`
}

type BalanceResponse struct {
	Balance string `json:"balance"`
	Utxoids []struct {
		Txid        string `json:"txID"`
		Outputindex int    `json:"outputIndex"`
	} `json:"utxoIDs"`
}

type GetUTXOsArgs struct {
	Addresses []string `json:"addresses"`
	Limit     int      `json:"limit"`
	Encoding  string   `json:"encoding"`
}

type GetUTXOsResponse struct {
	Numfetched string   `json:"numFetched"`
	Utxos      []string `json:"utxos"`
	Endindex   struct {
		Address string `json:"address"`
		Utxo    string `json:"utxo"`
	} `json:"endIndex"`
	Encoding string `json:"encoding"`
}

type ListAddressesResponse struct {
	Addresses []string `json:"addresses"`
}

type SendArgs struct {
	Assetid    string   `json:"assetID"`
	Amount     int      `json:"amount"`
	To         string   `json:"to"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Memo       string   `json:"memo"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type SendMultipleArgs struct {
	Outputs []struct {
		Assetid string `json:"assetID"`
		To      string `json:"to"`
		Amount  int    `json:"amount"`
	} `json:"outputs"`
	Memo       string   `json:"memo"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type SendNFTArgs struct {
	Assetid    string   `json:"assetID"`
	Groupid    int      `json:"groupID"`
	To         string   `json:"to"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type SendWalletArgs struct {
	Assetid    string   `json:"assetID"`
	Amount     int      `json:"amount"`
	To         string   `json:"to"`
	Memo       string   `json:"memo"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}
