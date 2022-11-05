# Gobase

⚠️ This is not production-ready software yet ⚠️
## Base settings tool for Golang services
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/2mikeg/gobase/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/2mikeg/gobase/tree/main) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


Gobase help to read and convert environment variables in its initially declared types to parse into native structs. ✨

### Download and install

    go get github.com/2mikeg/gobase@latest

### Features
You can read enviroment variables written in snakecase even if you write strucs fields as camelcase.

```go
type Settings struct {
	ThisVariable *int
} //ThisVariable field will be search as "THIS_VARIABLE"
```

### Example

```go
type Settings struct {
	VarOne string
	VarTwo *int //Pointers are supported
}

func main() {
	var s Settings
	sBytes, err := gobase.BaseModel(s)
	//BaseModel return a []byte that must be unmarshal in the struct

	if err != nil {
		panic(err)
	}

	json.Unmarshal(sBytes, &s) //
}
```

### Types
| Type       | Support |
|------------|---------|
| int8       | Yes     |
| int16      | Yes     |
| int32      | Yes     |
| int64      | Yes     |
| uint8      | No      |
| uint16     | No      |
| uint32     | Yes     |
| uint64     | Yes     |
| int        | Yes     |
| uint       | No      |
| rune       | Yes     |
| byte       | No      |
| uintptr    | No      |
| string     | Yes     |
| bool       | Yes     |
| complex64  | No      |
| complex128 | No      |

### Run tests

```sh
source scripts/test.sh
```