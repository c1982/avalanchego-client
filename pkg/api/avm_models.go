package api

//FixedCapAsset CreateFixedCapAsset	signature
type FixedCapAsset struct {
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

type AvmMint struct {
	Amount     int64    `json:"amount"`
	Assetid    string   `json:"assetID"`
	To         string   `json:"to"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type CreateVariableCapAsset struct {
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
