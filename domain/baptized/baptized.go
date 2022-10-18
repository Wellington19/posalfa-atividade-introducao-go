package baptized

import (
	"encoding/json"
	"fmt"

	"github.com/Wellington19/posalfa-atividade-introducao-go/domain"
)

type Baptized struct {
	TotalMeat               int `json:"total_meat"`
	TotalPeoples            int `json:"total_peoples"`
	TotalAccompaniment      int `json:"total_accompaniment"`
	TotalNonAlcoholicDrinks int `json:"total_nonalcoholic_drinks"`
}

func (b Baptized) ToJSON() ([]byte, error) {
	return json.Marshal(b)
}

type Service struct{}

func NewBaptized() Service {
	return Service{}
}

func (s Service) Calculate(p domain.Parameters) (domain.Result, error) {
	b := Baptized{}
	if p.NumberOfMens <= 0 || p.NumberOfWomen <= 0 || p.NumberOfChildrens <= 0 {
		return b, fmt.Errorf("must be informed the number of men, women and children and 0 is not accepted")
	}

	b.TotalMeat = p.NumberOfWomen*150 + p.NumberOfMens*200 + p.NumberOfChildrens*100
	b.TotalPeoples = p.NumberOfWomen + p.NumberOfMens + p.NumberOfChildrens
	b.TotalAccompaniment = b.TotalPeoples * 50
	b.TotalNonAlcoholicDrinks = 400 * b.TotalPeoples

	return b, nil
}
