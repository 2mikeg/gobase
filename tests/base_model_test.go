package tests

import (
	"testing"

	"github.com/2mikeg/gobase"
	"github.com/go-playground/assert/v2"
)

func TestBaseModel(t *testing.T) {

	type S struct {
		FirstVar  string
		SecondVar int
		ThirdVar  *bool
		FourthVar float64
	}

	var s S

	got := gobase.BaseModel(s)
	expected := map[string]interface{}{
		"FirstVar":  "test",
		"SecondVar": 1,
		"ThirdVar":  false,
		"FourthVar": 3.14,
	}

	assert.Equal(t, got, expected)

}
