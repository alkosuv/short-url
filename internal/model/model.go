package model

type URL struct {
	ID       int    `json:"id"`
	Original string `json:"original"`
	Hash     string `json:"hash"`
}
