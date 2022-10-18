package domain

type Parameters struct {
	NumberOfMens      int  `json:"number_of_mens"`
	NumberOfWomen     int  `json:"number_of_women"`
	NumberOfChildrens int  `json:"number_of_childrens"`
	HasAccompaniment  bool `json:"has_accompaniment"`
}

type Party interface {
	Calculate(Parameters) (Result, error)
}

type Result interface {
	ToJSON() ([]byte, error)
}
