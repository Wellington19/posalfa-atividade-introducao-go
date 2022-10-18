package barbecue

import (
	"encoding/json"
	"errors"

	"github.com/Wellington19/posalfa-atividade-introducao-go/domain"
)

type Barbecue struct {
	TotalMeat               int `json:"total_meat"`
	TotalPeoples            int `json:"total_peoples"`
	TotalAccompaniment      int `json:"total_accompaniment"`
	TotalNonAlcoholicDrinks int `json:"total_nonalcoholic_drinks"`
	TotalAlcoholicDrinks    int `json:"total_alcoholic_drinks"`
}

func (b Barbecue) ToJSON() ([]byte, error) {
	return json.Marshal(b)
}

type Service struct{}

func NewBarbecue() Service {
	return Service{}
}

func (s Service) Calculate(p domain.Parameters) (domain.Result, error) {
	b := Barbecue{}
	if p.NumberOfMens == 0 || p.NumberOfWomen == 0 {
		return b, errors.New("number of men and women must be greater than zero")
	}

	b.TotalMeat = p.NumberOfWomen*150 + p.NumberOfMens*200 + p.NumberOfChildrens*100
	b.TotalPeoples = p.NumberOfMens + p.NumberOfWomen + p.NumberOfChildrens
	b.TotalAccompaniment = b.TotalPeoples * 50
	b.TotalNonAlcoholicDrinks = 400 * b.TotalPeoples
	b.TotalAlcoholicDrinks = 500 * (p.NumberOfMens + p.NumberOfWomen)
	return b, nil
}
