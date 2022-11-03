package pkg

import "github.com/gobeam/stringy"

func Case2SnakeUpper(caseString string) string {
	str := stringy.New(caseString)
	strSnake := str.SnakeCase()
	return strSnake.ToUpper()
}
