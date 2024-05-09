package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewHaircut(t *testing.T) {
	expectedDate, err := time.Parse(time.DateTime, "2006-01-02 15:00:00")
	if err != nil {
		t.Error(err)
	}
	haircut, err := NewHaircut("Joao", 30.00, "2006-01-02 15:00:00")
	assert.Nil(t, err)
	assert.NotNil(t, haircut)
	assert.NotEmpty(t, haircut.ID)
	assert.NotEmpty(t, haircut.Price)
	assert.NotEmpty(t, haircut.Date)
	assert.Equal(t, expectedDate, haircut.Date)
	assert.Equal(t, "Joao", haircut.NameClient)
	assert.Equal(t, 30.00, haircut.Price)
}
