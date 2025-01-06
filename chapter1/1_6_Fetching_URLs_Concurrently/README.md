# 1.6 Fetching URLs Concurrently
- 下一个程序 `fetchall` 会并发获取多个URL, fetchall 的版本丢弃了响应，但报告了每个响应的大小和耗时
```go
// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
```
- `goroutine` 是并发函数的执行, `channel` 允许goroutine通过指定类型传递值给另一个goroutine的交互机制
- main函数在goroutine中运行, 而go语句创建了额外的goroutine
- main函数使用make创建了一个字符串类型的管道
- 对于每个命令行参数，第一个 range 循环中的 go 语句启动一个新的 goroutine，异步调用 fetch 函数使用 http.Get 获取 URL
- 每当结果到达时，fetch 就会在通道 ch 上发送一条摘要行
- 当一个goroutine尝试在管道发送或接收，他会一直阻塞直到其他goroutine尝试对应的接收或发送操作
  - 此时，值被传输，两个goroutine都继续执行
- 