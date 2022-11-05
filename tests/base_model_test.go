package tests

import (
	"encoding/json"
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

	got, bsErr := gobase.BaseModel(s)

	if bsErr != nil {
		panic(bsErr)
	}

	expectedMap := map[string]interface{}{
		"FirstVar":  "test",
		"SecondVar": 1,
		"ThirdVar":  false,
		"FourthVar": 3.14,
	}

	expected, err := json.Marshal(expectedMap)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, got, expected)

}
