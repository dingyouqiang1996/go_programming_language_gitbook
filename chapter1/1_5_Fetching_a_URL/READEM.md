# 1.5 Fetching a URL
- Go使用 `net` 包提供访问网络的的功能
- ，这里有一个fetch程序，它获取每个指定URL的内容并将其作为未经解释的文本打印出来；它受到了非常有用的工具curl的启发
- `fetch.go`
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
```
  - 这个程序会创建一个HTTP请求, 如果没有发生错误则会返回一个 `resp` 结构体
  - `resp` 结构体的 `Body` 字段包含了服务器回复而来的可读的流
  - 然后 `ioutil.ReadAll` 会读取整段的流,把结果存储在 `b`
  - `resp.Body.Close` 会关闭流, 防止资源泄露

