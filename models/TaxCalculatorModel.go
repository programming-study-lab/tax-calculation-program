package models

type TaxCalculator struct {
	TotalIncome int          `json:"totalIncome"`
	Wht         int          `json:"wht"`
	Allowances  []Allowances `json:"allowances"`
}

type Allowances struct {
	AllowancesType string `json:"allowancesType"`
	Amount         int    `json:"amount"`
}

type TaxCalculatorResponse struct {
	Tax      int        `json:"tax"`
	TaxLevel []TaxLevel `json:"taxLevel"`
}

type TaxLevel struct {
	Level string `json:"level"`
	Tax   int    `json:"tax"`
}
