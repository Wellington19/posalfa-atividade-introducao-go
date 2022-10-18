package barbecue

import (
	"testing"

	"github.com/Wellington19/posalfa-atividade-introducao-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBarbecueGhost(t *testing.T) {
	s := NewBarbecue()
	_, err := s.Calculate(domain.Parameters{
		NumberOfMens:      0,
		NumberOfWomen:     1,
		NumberOfChildrens: 2,
		HasAccompaniment:  true,
	})
	assert.Equal(t, err.Error(), "number of men and women must be greater than zero")
}

func TestCalculateBarbecuePumping(t *testing.T) {
	s := NewBarbecue()
	ch, err := s.Calculate(domain.Parameters{
		NumberOfMens:      3,
		NumberOfWomen:     3,
		NumberOfChildrens: 6,
		HasAccompaniment:  true,
	})
	assert.Nil(t, err)
	expected := Barbecue{
		TotalMeat:               1650,
		TotalPeoples:            12,
		TotalAccompaniment:      600,
		TotalNonAlcoholicDrinks: 4800,
		TotalAlcoholicDrinks:    3000,
	}
	assert.Equal(t, ch, expected)
}
