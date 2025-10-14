package model

type Relatory struct {
	Data []RelatoryItem `json:"data"`
	Page int            `json:"page"`
}

type RelatoryItem struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
}
