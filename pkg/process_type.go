package pkg

import "strconv"

func checkPointer(envVarTypeRaw string) string {
	firstCharacter := string(envVarTypeRaw[0])

	var envVarType string

	if firstCharacter == "*" {
		envVarType = envVarTypeRaw[1:]
	} else {
		envVarType = envVarTypeRaw
	}

	return envVarType
}

func ProcessType(envVar, envVarTypeRaw string) (interface{}, error) {

	var processVariable interface{}
	var err error

	envVarType := checkPointer(envVarTypeRaw)

	switch envVarType {

	case "bool":
		processVariable, err = strconv.ParseBool(envVar)

	case "int":
		processVariable, err = strconv.Atoi(envVar)

	case "int8":
		processVariable, err = strconv.ParseInt(envVar, 10, 8)

	case "int16":
		processVariable, err = strconv.ParseInt(envVar, 10, 16)

	case "int32", "rune":
		processVariable, err = strconv.ParseInt(envVar, 10, 32)

	case "int64":
		processVariable, err = strconv.ParseInt(envVar, 10, 64)

	case "float32":
		processVariable, err = strconv.ParseFloat(envVar, 32)

	case "float64":
		processVariable, err = strconv.ParseFloat(envVar, 64)

	case "string":
		processVariable, err = envVar, nil
	}

	return processVariable, err

}
