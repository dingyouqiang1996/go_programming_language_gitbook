# 2.4 Assignments
- 变量持有的值通过赋值语句更新，其最简单的形式是在 = 符号的左侧有一个变量，在右侧有一个表达式
```go
x = 1 // named variable
*p = true // indirect variable
person.anme = "bob" // struct field
count[x] = count[x] * scale // array or slice or map element
```
- 每个算术和位运算的二元运算符都有一个对应的赋值运算符，允许例如将最后一条语句重写为
```go
count[x] *= scale
```
- 这使我们不必重复（和重新计算）变量的表达式。数值变量还可以通过 ++ 和 -- 语句进行递增和递减
```go
v := 1
v++ // same as v = v + 1; v becomes 2
v-- // same as v = v - 1; v becomes 1 again
```
## 2.4.1 Tuple Assignment
- 另一种赋值形式称为元组赋值，允许同时为多个变量赋值。所有右手边的表达式在任何变量更新之前都被计算，这使得这种形式在一些变量出现在赋值的两边时最有用，例如在交换两个变量的值时
```go
x, y = y, x
a[i], a[j] = a[j], a[i]
```
- 或者在计算两个整数的最大公约数（GCD）时
```go
func fib(n int) int {
      x, y := 0, 1
      for i := 0; i < n; i++ {
          x, y = y, x+y
      }
      return x
}
```
- 或者在迭代计算第 n 个斐波那契数时。
```go
func fib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x
}
```
- 元组赋值还可以使一系列简单的赋值更加紧凑
```go
i, j, k = 2, 3, 5
```
- 某些表达式，例如调用返回多个结果的函数，会产生多个值。当这样的调用在赋值语句中使用时，左手边必须有与函数结果数量相同的变量。
```go
f, err = os.Open("foo.txt") // fucntion call returns two values
```
- 通常，函数使用这些额外的结果来指示某种错误，要么通过返回一个错误，如 os.Open 的调用，要么返回一个布尔值，通常称为 ok。正如我们将在后面的章节中看到的，还有三个运算符有时也表现出这种方式。如果在期望有两个结果的赋值中出现映射查找、类型断言或通道接收，每个都会产生一个额外的布尔结果
```go
v, ok = m[key] // map lookup
v, ok = x.(T) // type assertion
v, ok = <- ch // channel receive
```
- 与变量声明一样，我们可以将不需要的值赋给空白标识符
```go
_, err = io.Copy(dst, src) // discard byte count
_, ok = x.(T) // check type but discard result
```
## 2.4.2 Assignability
- 赋值语句是一种显式的赋值形式，但在程序的许多地方会发生隐式赋值：
  - 函数调用隐式地将参数值赋给相应的参数变量
  - 返回语句隐式地将返回操作数赋给相应的结果变量
  - 并且像这个切片这样的复合类型的字面量表达式隐式地将值赋给切片的元素
```go
medals := []string{"gold", "silver", "bronze"}
```
- 隐式地为每个元素赋值，就好像它是这样写的
```go
medals[0] = "gold"
medals[1] = "silver"
medals[2] = "bronze"
```
- 尽管映射和通道的元素不是普通变量，但它们也受到类似的隐式赋值的影响
- 如果左手边（变量）和右手边（值）具有相同的类型，则显式或隐式的赋值总是合法的。更一般地说，只有当值可以赋给变量的类型时，赋值才是合法的
- 两个值是否可以用 == 和 != 进行比较与可赋值性有关：在任何比较中，第一个操作数必须可以赋给第二个操作数的类型，或者反之亦然。与可赋值性一样，我们将在介绍每种新类型时解释相关情况

