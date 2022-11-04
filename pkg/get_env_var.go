package pkg

import (
	"log"
	"os"
)

func GetEnvVar(fields []fieldProperties) map[string]interface{} {
	var baseSettingsMap = make(map[string]interface{})

	for f := range fields {

		EnvVar := os.Getenv(fields[f].SnakeName)

		if EnvVar != "" {
			EnvVarType := fields[f].Type
			processVariable, err := ProcessType(EnvVar, EnvVarType)

			if err != nil {
				log.Panic(err)
			}

			baseSettingsMap[fields[f].Name] = processVariable
		}

	}

	return baseSettingsMap
}
