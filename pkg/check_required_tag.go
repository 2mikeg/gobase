package pkg

import "errors"

func checkRequired(fieldTag string) (bool, error) {
	optionalTag := `bs:"optional"`
	if fieldTag == optionalTag {
		return false, nil
	} else if fieldTag == "" {
		return true, nil
	} else {
		return false, errors.New("not implemented tag")
	}
}
