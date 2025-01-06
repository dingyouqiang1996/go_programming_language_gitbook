# 1.7 A Web Server
- Go 的库使编写响应客户端请求（如 fetch 提出的请求）的网络变得简单易行
- 在本节中，我们将展示一个返回 URL 路径组件的最小服务器
  - `http://localhost:8000/hello` 
  - 这个程序只有几行长，因为库函数做了大部分工作。
  - 主函数将一个处理函数连接到以 / 开头的所有传入 URL，并启动一个服务器，监听端口 8000 上的传入请求
  - 请求被表示为一个类型为 http.Request 的结构体，其中包含许多相关字段，其中之一是传入请求的 URL。
  - 当请求到达时，它会被传递给处理函数，该函数从请求 URL 中提取路径部分（例如 /hello），并使用 fmt.Fprintf 将其作为响应返回
- 一个有用的补充是提供一个返回某种状态的特定 URL。例如，这个版本除了进行相同的回显外，还统计了请求的数量；对 URL /count 的请求会返回到目前为止的计数，不包括对 /count 本身的请求
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requests URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
```
- 服务器有两个处理函数，请求 URL 决定调用哪一个：对 /count 的请求会调用 counter，其他所有请求会调用 handler。以斜杠结尾的处理模式会匹配任何以该模式为前缀的 URL
- 在幕后，服务器为每个传入请求在单独的 goroutine 中运行处理函数，以便可以同时处理多个请求
- 然而，如果两个并发请求同时尝试更新 count，它可能无法一致地递增；程序会有一个严重的错误，称为竞态条件
- 为了避免这个问题，我们必须确保最多只有一个 goroutine 同时访问该变量，这就是 mu.Lock() 和 mu.Unlock() 调用围绕每次访问 count 的目的
- 作为一个更丰富的例子，处理函数可以报告它接收到的头部和表单数据，使服务器可用于检查和调试请求
```go
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
```
- 在 Go 语言中，你可以在 if 语句的条件之前放置一个简单的语句，例如局部变量的声明。这在错误处理中特别有用，因为它允许你声明并初始化一个变量，然后立即检查错误。以下是一个典型的示例
```go
err := r.ParseForm()
if err != nil {
    log.Print(err)
}
```
- 但将这些语句组合在一起更简洁，并且减少了变量 `err` 的作用域，这是一种良好的实践
- 之前的程序用了三种输出: `os.Stdout`, `ioutil.Discard`, `http.Writer`，他们的使用细节不相同，但是他们都实现了 `io.Writer` 接口
- 比如将lissajous写入http输出
```go
hadnler := func(w http.ResponseWriter, r *http.Request) {
    lissajous(w)
}
http.HandleFunc("/", hadnler)
```
