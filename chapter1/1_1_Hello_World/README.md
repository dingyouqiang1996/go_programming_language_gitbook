# 1.1 Hello World
- `ch1/helloworld.go`
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```
- 运行
```shell
$ go run main.go
Hello, 世界
```
- Go是 `编译型语言`, Go工具链使用 `计算机原生机器语言` 转化 `源程序`
  - Go工具链包括 `run` 子命令
  - `源程序` 是以 `.go` 结尾的代码
  - 编译时会链接所需的 `库`
  - 运行结果是生成一个 `可执行的二进制文件`
  - Go 原生地处理 `Unicode` 万国码
  - Go 通过 `包` 组织代码
    - 使用 `import` 关键字导入相关的包
  - `main包` 的 `main func()函数` 是程序的入口
