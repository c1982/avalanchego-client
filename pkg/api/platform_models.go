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
