package logger2_test

import (
	"testing"

	"github.com/reddotpay/logger2"
	"github.com/stretchr/testify/assert"
)

func TestLogger_JSON_MaskNumber(t *testing.T) {
	s := `{"number":"4111112111115432"}`
	assert.Equal(t, "{\"number\":\"411111******5432\"}", logger2.MaskCard(s))
}
func TestLogger_JSON_MaskCVV(t *testing.T) {
	s := `{"cvv":"111"}`
	assert.Equal(t, "{\"cvv\":\"***\"}", logger2.MaskCard(s))
}

func TestLogger_JSON_MaskEmail(t *testing.T) {
	s := `{"email":"hello@domain.com"}`
	assert.Equal(t, "{\"email\":\"hel**@domain.com\"}", logger2.MaskCard(s))
}
