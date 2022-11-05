package gobase

import (
	"encoding/json"

	"github.com/2mikeg/gobase/pkg"
)

func BaseModel(s any) ([]byte, error) {

	structFields := pkg.StructChecker(s)
	envVariablesAsMap := pkg.GetEnvVar(structFields)

	b, err := json.Marshal(envVariablesAsMap)

	return b, err

}
