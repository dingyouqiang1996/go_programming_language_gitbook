# 1.3 查找重复的行
- 这里会展示三种 `dup` 的程序, 灵感来自unix命令 `uniq`
- dup会打印标准输入计数超过一行的行
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
  - go的if条件部分不需要 `()`, 但是语句块需要用 `{}` 扩住
    - if 条件返回false可执行else部分
  - map提供常数时间复杂度的操作来存储、检索或测试集合中的项
    - map的键是任意可以用 == 比较的类型
	- make(map[string]int) 评估右边的零值
	- map的迭代顺序实际上是随机的
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
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
// NOTE: ignoring potential errors from input.Err()
```
  - 上面代码以流模式遍历读取文件参数, 使用模块进行行数读取然后打印
- 第三个版本会被简化, 首先只读取命名文件, 而非标准输入
```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
```
  - ReadFile 会返回一个字节切片, 里面的内容必须通过 string() 转换类型
- 很少会使用 os.File 这种低等级的文件操作, 而是使用 bufio、io/util 包中的更高级的函数