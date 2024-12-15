# goast
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/mtresnik/goast/blob/main/LICENSE)
[![version](https://img.shields.io/badge/version-1.1.1-blue)](https://github.com/mtresnik/goast/releases/tag/v1.1.1)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-green.svg?style=flat-square)](https://makeapullrequest.com)
<hr>

Goast (pronounced Ghost) is a Go implementation of an AST and parser. This allows for strings to be converted to and from mathematical structures.


### Sample Code

In your project run:
```
go mod download github.com/mtresnik/goutils
go mod download github.com/mtresnik/goast 
```

Your `go.mod` file should look like this:
```go 
module mymodule

go 1.23.3

require github.com/mtresnik/goast v1.1.1
```


Then in your go files you should be able to access the parser:
```go 
package main

import (
	"github.com/mtresnik/goast/pkg/goast"
)

func main() {
	operation, err := goast.ParseOperation("a * bc + 123 / sin(3.1415 * n) ^ log_(2, 8) - e")
	if err != nil {
		println((*err).Error())
		return
	}
	println(operation.String())
}
```