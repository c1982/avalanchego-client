package api

type AddDelegatorArgs struct {
	Nodeid        string   `json:"nodeID"`
	Rewardaddress string   `json:"rewardAddress"`
	Starttime     int      `json:"startTime"`
	Endtime       int      `json:"endTime"`
	Stakeamount   int      `json:"stakeAmount"`
	From          []string `json:"from"`
	Changeaddr    string   `json:"changeAddr"`
	Username      string   `json:"username"`
	Password      string   `json:"password"`
}

type AddValidatorArgs struct {
	Nodeid            string   `json:"nodeID"`
	Rewardaddress     string   `json:"rewardAddress"`
	From              []string `json:"from"`
	Changeaddr        string   `json:"changeAddr"`
	Starttime         string   `json:"startTime"`
	Endtime           string   `json:"endTime"`
	Stakeamount       int      `json:"stakeAmount"`
	Delegationfeerate int      `json:"delegationFeeRate"`
	Username          string   `json:"username"`
	Password          string   `json:"password"`
}

type AddSubnetValidatorArgs struct {
	Nodeid     string   `json:"nodeID"`
	Subnetid   string   `json:"subnetID"`
	Starttime  int      `json:"startTime"`
	Endtime    int      `json:"endTime"`
	Weight     int      `json:"weight"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type CreateBlockchainArgs struct {
	Vmid        string   `json:"vmID"`
	Subnetid    string   `json:"subnetID"`
	Name        string   `json:"name"`
	Genesisdata string   `json:"genesisData"`
	Encoding    string   `json:"encoding"`
	From        []string `json:"from"`
	Changeaddr  string   `json:"changeAddr"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
}

type CreateSubnetArgs struct {
	Controlkeys []string `json:"controlKeys"`
	Threshold   int      `json:"threshold"`
	From        []string `json:"from"`
	Changeaddr  string   `json:"changeAddr"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
}

type PlatformExportAVAXArgs struct {
	To         string   `json:"to"`
	Amount     int      `json:"amount"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

type GetBalanceResponse struct {
	Balance            string `json:"balance"`
	Unlocked           string `json:"unlocked"`
	Lockedstakeable    string `json:"lockedStakeable"`
	Lockednotstakeable string `json:"lockedNotStakeable"`
	Utxoids            []struct {
		Txid        string `json:"txID"`
		Outputindex int    `json:"outputIndex"`
	} `json:"utxoIDs"`
}

type GetBlockchainsResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Blockchains []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Subnetid string `json:"subnetID"`
			Vmid     string `json:"vmID"`
		} `json:"blockchains"`
	} `json:"result"`
	ID int `json:"id"`
}

type GetCurrentValidatorsResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Validators []struct {
			Txid        string `json:"txID"`
			Starttime   string `json:"startTime"`
			Endtime     string `json:"endTime"`
			Stakeamount string `json:"stakeAmount"`
			Nodeid      string `json:"nodeID"`
			Rewardowner struct {
				Locktime  string   `json:"locktime"`
				Threshold string   `json:"threshold"`
				Addresses []string `json:"addresses"`
			} `json:"rewardOwner"`
			Potentialreward string `json:"potentialReward"`
			Delegationfee   string `json:"delegationFee"`
			Uptime          string `json:"uptime"`
			Connected       bool   `json:"connected"`
			Delegators      []struct {
				Txid        string `json:"txID"`
				Starttime   string `json:"startTime"`
				Endtime     string `json:"endTime"`
				Stakeamount string `json:"stakeAmount"`
				Nodeid      string `json:"nodeID"`
				Rewardowner struct {
					Locktime  string   `json:"locktime"`
					Threshold string   `json:"threshold"`
					Addresses []string `json:"addresses"`
				} `json:"rewardOwner"`
				Potentialreward string `json:"potentialReward"`
			} `json:"delegators"`
		} `json:"validators"`
	} `json:"result"`
	ID int `json:"id"`
}

type GetPendingValidatorsResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Validators []struct {
			Txid          string `json:"txID"`
			Starttime     string `json:"startTime"`
			Endtime       string `json:"endTime"`
			Stakeamount   string `json:"stakeAmount"`
			Nodeid        string `json:"nodeID"`
			Delegationfee string `json:"delegationFee"`
			Connected     bool   `json:"connected"`
		} `json:"validators"`
		Delegators []struct {
			Txid        string `json:"txID"`
			Starttime   string `json:"startTime"`
			Endtime     string `json:"endTime"`
			Stakeamount string `json:"stakeAmount"`
			Nodeid      string `json:"nodeID"`
		} `json:"delegators"`
	} `json:"result"`
	ID int `json:"id"`
}

type GetStakingAssetIDResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Subnets []struct {
			ID          string   `json:"id"`
			Controlkeys []string `json:"controlKeys"`
			Threshold   string   `json:"threshold"`
		} `json:"subnets"`
	} `json:"result"`
	ID int `json:"id"`
}

type GetSubnetsResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Subnets []struct {
			ID          string   `json:"id"`
			Controlkeys []string `json:"controlKeys"`
			Threshold   string   `json:"threshold"`
		} `json:"subnets"`
	} `json:"result"`
	ID int `json:"id"`
}

type GetStakeArgs struct {
	Addresses []string `json:"addresses"`
}

type PlatformGetUTXOsArgs struct {
	Addresses []string `json:"addresses"`
	Limit     int      `json:"limit"`
	Encoding  string   `json:"encoding"`
}

type PlatformGetUTXOsResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Numfetched string   `json:"numFetched"`
		Utxos      []string `json:"utxos"`
		Endindex   struct {
			Address string `json:"address"`
			Utxo    string `json:"utxo"`
		} `json:"endIndex"`
		Encoding string `json:"encoding"`
	} `json:"result"`
	ID int `json:"id"`
}
