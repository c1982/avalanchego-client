package api

type GetLastAcceptedResponse struct {
	ID        string `json:"id"`
	Bytes     string `json:"bytes"`
	Timestamp string `json:"timestamp"`
	Encoding  string `json:"encoding"`
	Index     string `json:"index"`
}

type GetContainerByIndexResponse struct {
	ID        string `json:"id"`
	Bytes     string `json:"bytes"`
	Timestamp string `json:"timestamp"`
	Encoding  string `json:"encoding"`
	Index     string `json:"index"`
}

type GetContainerRangeResponse struct {
	ID        string `json:"id"`
	Bytes     string `json:"bytes"`
	Timestamp string `json:"timestamp"`
	Encoding  string `json:"encoding"`
	Index     string `json:"index"`
}
