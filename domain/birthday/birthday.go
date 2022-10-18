package birthday

import (
	"encoding/json"
	"fmt"

	"github.com/Wellington19/posalfa-atividade-introducao-go/domain"
)

type Birthday struct {
	TotalSavouries   int `json:"total_savouries"`
	TotalCake        int `json:"total_cake"`
	TotalSweet       int `json:"total_sweet"`
	TotalRefrigerant int `json:"total_refrigerant"`
	TotalPeoples     int `json:"total_peoples"`
}

func (b Birthday) ToJSON() ([]byte, error) {
	return json.Marshal(b)
}

type Service struct{}

func NewBirthday() Service {
	return Service{}
}

func (s Service) Calculate(p domain.Parameters) (domain.Result, error) {
	b := Birthday{}
	if p.NumberOfMens <= 0 || p.NumberOfWomen <= 0 || p.NumberOfChildrens <= 0 {
		return b, fmt.Errorf("must be informed the number of men, women and children and 0 is not accepted")
	}

	totalAdults := p.NumberOfMens + p.NumberOfWomen
	b.TotalPeoples = totalAdults + p.NumberOfChildrens
	b.TotalSavouries = p.NumberOfChildrens*6 + totalAdults*10
	b.TotalCake = b.TotalPeoples * 100
	b.TotalSweet = b.TotalPeoples * 4
	b.TotalRefrigerant = b.TotalPeoples * 500

	return b, nil
}
