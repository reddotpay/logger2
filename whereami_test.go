package logger2_test

import (
	"strings"
	"testing"

	"github.com/reddotpay/logger2"
	"github.com/stretchr/testify/assert"
)

func Test_WhereAmI(t *testing.T) {
	caller := logger2.WhereAmI()
	assert.True(t, strings.Contains(caller, "github.com/reddotpay/logger2_test.Test_WhereAmI"))
}
