// Code generated by sqlc. DO NOT EDIT.

package postgres

import ()

type Article struct {
	Harvestid    string `json:"harvestid"`
	Cerebroscore string `json:"cerebroscore"`
	Url          string `json:"url"`
	Title        string `json:"title"`
	Cleanimage   string `json:"cleanimage"`
}

type Contentmarketing struct {
	Harvestid         string `json:"harvestid"`
	Commercialpartner string `json:"commercialpartner"`
	Logourl           string `json:"logourl"`
	Cerebroscore      string `json:"cerebroscore"`
}
