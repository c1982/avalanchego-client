package api

type NewTokenArgs struct {
	Password  string   `json:"password"`
	Endpoints []string `json:"endpoints"`
}
