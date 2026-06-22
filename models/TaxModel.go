package models

type TaxCalculator struct {
	TotalIncome int `json:"totalIncome"`
	Wht         int `json:"wht"`
	Allowances  []Allowances
}

type Allowances struct {
	AllowancesType string `json: "allowancesType"`
	Amount         int    `json: "amount"`
}

// var Allwoances = []AllowancesPayload{}
