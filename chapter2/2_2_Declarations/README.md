# 2.2 Declarations
- 声明为程序实体命名并指定其部分或全部属性。有四种主要的声明类型：var、const、type 和 func
- Go 程序存储在一个或多个以 .go 结尾的文件中。每个文件都以一个包声明开始，说明该文件属于哪个包。包声明之后是任何导入声明，然后是类型、变量、常量和函数的包级声明序列，顺序任意
- `boiling.go`
```go
// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}
```


