package model

type Address struct {
	Id       int    `json:"id"`
	City     string `json:"city"`
	State    string `json:"state"`
	Street1  string `json:"street1"`
	Street2  string `json:"street2"`
	Zip_code int    `json:"zip_code"`
}
