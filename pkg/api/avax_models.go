package api

type AVAXGetUTXOsArgs struct {
	Addresses   []string `json:"addresses"`
	Sourcechain string   `json:"sourceChain"`
	Startindex  struct {
		Address string `json:"address"`
		Utxo    string `json:"utxo"`
	} `json:"startIndex"`
	Encoding string `json:"encoding"`
}
type AVAXGetUTXOsResponse struct {
	Numfetched string   `json:"numFetched"`
	Utxos      []string `json:"utxos"`
	Endindex   struct {
		Address string `json:"address"`
		Utxo    string `json:"utxo"`
	} `json:"endIndex"`
	Encoding string `json:"encoding"`
}
