package model

type ExternalRelatory struct {
	Data []ExternalRelatoryItem `json:"data"`
	Page int                    `json:"page"`
}

type ExternalRelatoryItem struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
}
