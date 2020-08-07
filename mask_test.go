package logger2_test

import (
	"fmt"
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

func TestLoggerN(t *testing.T) {
	s := `OrderRef1=d41d8cd98f00b204e9800998ecf8427e&OrderRef2=9fd3499cad174fa7e895b8e84b6d2009&amount=5000.000000&cardHolder=&cardNo=&currCode=764&epMonth=0&epYear=0&lang=E&merchantId=REDDOT&orderRef=627988e4c6684602a88980a658fc0745&pMethod=&payType=H&vbvTransaction=F`

	fmt.Println(logger2.MaskCard(s))

}
func TestLoggerPassword(t *testing.T) {
	s := `password=12345`

	fmt.Println(logger2.MaskCard(s))

}
