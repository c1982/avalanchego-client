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
