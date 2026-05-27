package main

type Company struct {
	ID            int    `json:"id"`
	CompanyName   string `json:"company_name"`
	INN           string `json:"inn"`
	KPP           string `json:"kpp"`
	OGRN          string `json:"ogrn"`
	OKPO          string `json:"okpo"`
	LegalForm     string `json:"legal_form"`
	LegalAddress  string `json:"legal_address"`
	Status        string `json:"status"`
	InclusionDate string `json:"inclusion_date"`
}

type CompanyFilter struct {
	CompanyName  string `json:"company_name"`
	LegalForm    string `json:"legal_form"`
	LegalAddress string `json:"legal_address"`
	Status       string `json:"status"`
}
