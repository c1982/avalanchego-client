package api

import "time"

type InfoPeersArgs struct {
	Numpeers int `json:"numPeers"`
	Peers    []struct {
		IP           string    `json:"ip"`
		Publicip     string    `json:"publicIP"`
		Nodeid       string    `json:"nodeID"`
		Version      string    `json:"version"`
		Lastsent     time.Time `json:"lastSent"`
		Lastreceived time.Time `json:"lastReceived"`
	} `json:"peers"`
}
