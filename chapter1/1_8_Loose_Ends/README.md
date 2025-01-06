# 1.8 Loose Ends
- switch 多路分支
```go
switch coinflip() {
case "heads":
    heads++
case "tails":
    tails++
default:
    fmt.Println("landed on edge!")
}
```
- 调用 coinflip 的结果与每个 case 的值进行比较。case 从上到下进行评估，因此第一个匹配的 case 将被执行
- switch 不需要操作数；它只需列出 case，每个 case 都是一个布尔表达式
```go
func Signum(x int) int {
    switch {
    case x > 0:
        return +1
    default:
        return 0
    case x < 0:
        return -1
    }
}
```
- 这种形式被称为无标签 switch；它等同于 switch true。
- 命名类型：类型声明使得可以给现有类型命名。由于结构体类型通常很长，它们几乎总是被命名。一个熟悉的例子是为 2-D 图形系统定义一个 Point 类型：
```go
type Point struct {
    X, Y int
}
var p Point
```
- 指针：Go 提供了指针，即包含变量地址的值
- 



