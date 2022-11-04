package gobase

import "github.com/2mikeg/gobase/pkg"

func BaseModel(s any) map[string]interface{} {

	structFields := pkg.StructChecker(s)
	envVariablesAsMap := pkg.GetEnvVar(structFields)

	return envVariablesAsMap

}
