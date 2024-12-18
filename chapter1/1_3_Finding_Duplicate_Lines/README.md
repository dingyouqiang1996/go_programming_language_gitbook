# 1.3 查找重复的行
- 这里会展示三种 `dup` 的程序, 灵感来自unix命令 `uniq`
- dup的第一个版本会打印标准输入中出现超过一次的每一行，并在其前面加上它的计数
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
```
  - `bufio.NewScanner(os.Stdin)` 会创建一个扫描器, 从标准输入读取内容
  - 扫描器从程序的标准输入中读取数据。
    - 每次调用input.Scan()会读取下一行，并从末尾移除换行字符
    - 通过调用input.Text()可以获取结果
    - Scan函数在还有行可读时返回true，在没有更多输入时返回false
  - `fmt.Printf` 格式化的函数
  ```
  %d 十进制整数
  %x, %o, %b 整数的十六进制，八进制，二进制表示
  %f, %g, %e 浮点数： 3.141593 3.141592653589793 3.141593e+00
  %t 布尔：true或false
  %c 字符（rune） (Unicode码点)
  %s 字符串
  %q 字符串或字节切片的带引号的go语法表示
  %v 值的默认格式表示
  %T 值的类型的go语法表示
  %% 字面上的百分号，并非值的占位符
  ```
- 第二个版本读取的文件名作为参数, 并打印文件中重复的行
