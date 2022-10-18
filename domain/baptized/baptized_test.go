package baptized

import (
	"testing"

	"github.com/Wellington19/posalfa-atividade-introducao-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBaptizedIncomplete(t *testing.T) {
	s := NewBaptized()
	_, err := s.Calculate(domain.Parameters{
		NumberOfMens:      0,
		NumberOfWomen:     1,
		NumberOfChildrens: 2,
		HasAccompaniment:  true,
	})
	assert.Equal(t, err.Error(), "must be informed the number of men, women and children and 0 is not accepted")
}

func TestCalculateBaptizedPumping(t *testing.T) {
	s := NewBaptized()
	ch, err := s.Calculate(domain.Parameters{
		NumberOfMens:      3,
		NumberOfWomen:     3,
		NumberOfChildrens: 6,
		HasAccompaniment:  true,
	})
	assert.Nil(t, err)
	expected := Baptized{
		TotalMeat:               1650,
		TotalPeoples:            12,
		TotalAccompaniment:      600,
		TotalNonAlcoholicDrinks: 4800,
	}
	assert.Equal(t, ch, expected)
}
