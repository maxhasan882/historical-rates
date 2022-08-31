package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToDate(t *testing.T) {
	date := "2006-01-01"
	time, err := StringToDate(date, "2006-01-01")
	assert.Equal(t, err, nil)
	assert.Equal(t, time.Format("2006-01-01"), date)
}
