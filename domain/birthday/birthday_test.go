package birthday

import (
	"testing"

	"github.com/Wellington19/posalfa-atividade-introducao-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBirthdayGhost(t *testing.T) {
	s := NewBirthday()
	_, err := s.Calculate(domain.Parameters{
		NumberOfMens:      0,
		NumberOfWomen:     1,
		NumberOfChildrens: 6,
	})
	assert.Equal(t, err.Error(), "must be informed the number of men, women and children and 0 is not accepted")
}

func TestCalculateBirthdayPumping(t *testing.T) {
	s := NewBirthday()
	ch, err := s.Calculate(domain.Parameters{
		NumberOfMens:      3,
		NumberOfWomen:     3,
		NumberOfChildrens: 20,
	})
	assert.Nil(t, err)
	expected := Birthday{
		TotalSavouries:   180,
		TotalCake:        2600,
		TotalSweet:       104,
		TotalRefrigerant: 13000,
		TotalPeoples:     26,
	}
	assert.Equal(t, ch, expected)
}
