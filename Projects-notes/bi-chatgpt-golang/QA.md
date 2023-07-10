---
share: true  
---

# 这里总结初学Go语言的开发过程问题总结

`好记性不如烂文档`

# 1. 关于init函数和构造函数
## init函数

在Go语言中，`init`函数是一种特殊的函数，它可以在程序执行前自动被调用。每个包可以包含一个或多个`init`函数，它们会按照它们在源代码中的顺序被调用。

`init`函数没有参数和返回值，它的定义形式如下：

```go
func init() {
    // 初始化代码
}
```

`init`函数通常用于执行一些初始化操作，例如初始化全局变量、注册驱动程序、加载配置文件等。它在程序执行前被调用，因此适合用来进行一些必要的准备工作。注意，`init`函数不能被显式地调用，它由Go运行时自动调用。

## 构造函数

构造函数在Go语言中没有像其他语言那样的特殊语法，但可以通过惯例来模拟构造函数的行为。通常，我们会定义一个函数，以`New`开头，并返回一个结构体的指针。这个函数可以进行初始化操作，并返回一个已经初始化好的结构体实例。

下面是一个示例：

```go
type Person struct {
    Name string
    Age  int
}

func NewPerson(name string, age int) *Person {
    p := Person{
        Name: name,
        Age:  age,
    }
    return &p
}
```

在上面的示例中，我们定义了一个`Person`结构体和一个`NewPerson`函数作为构造函数。`NewPerson`函数接收姓名和年龄作为参数，并返回一个已经初始化好的`Person`结构体指针。

构造函数的使用场景包括但不限于以下几种情况：
- 创建一个复杂的对象实例时，构造函数可以隐藏创建过程的细节，提供一个简单的接口。
- 进行一些必要的初始化操作，例如设置默认值、注册回调函数等。
- 封装一些特定的逻辑，确保对象的创建过程符合某种规范。

总而言之，`init`函数用于包的初始化，而构造函数用于创建对象实例并进行必要的初始化操作。它们在Go语言中有着不同的用途和规范。





# 2. 关于Go语言中的&和\*的运用场景以及知识点

  在Go语言中，`&`和`*`是两个重要的运算符，用于处理指针和地址相关的操作。

1. `&`运算符：
   它用于获取变量的地址。例如，`&x`将返回变量`x`的地址。可以将地址赋值给指针变量，用于后续的操作。

示例代码：

```go
package main

import "fmt"

func main() {
    x := 10
    fmt.Println(&x) // 输出变量x的地址
}
```

2. `*`运算符：
它用于解引用指针，即获取指针指向的值。例如，`*p`将返回指针`p`所指向的值。

示例代码：

```go
package main

import "fmt"

func main() {
    x := 10
    p := &x       // 获取变量x的地址
    fmt.Println(*p) // 输出指针p所指向的值
}
```

`&`和`*`经常一起使用，用于传递指针或引用，以及在函数中修改变量的值。

示例代码：

```go
package main

import "fmt"

func changeValue(x *int) {
    *x = 20 // 修改指针所指向的值
}

func main() {
    x := 10
    p := &x         // 获取变量x的地址
    changeValue(p)  // 传递指针给函数
    fmt.Println(x) // 输出修改后的值
}
```

这是关于`&`和`*`的一些基本知识点和运用场景。它们在Go语言中常用于处理指针和地址相关的操作，可以用于传递指针、修改变量的值等。




# 3. 变量的首字母大写和小写
1. 首字母大写的变量：这种变量被称为**导出变量**（Exported Variable），可以被其他包（package）访问和使用。在Go语言中，变量名的可见性是由首字母的大小写来决定的。如果一个变量的首字母大写，表示该变量是公开的，可以在其他包中被引用和使用。

示例：
```go
package main  
import "fmt"  // 首字母大写的变量可以被其他包访问 

var PublicVariable string = "Hello"  
func main() {
	fmt.Println(PublicVariable) 
}
```


2. 首字母小写的变量：这种变量被称为**非导出变量**（Unexported Variable），只能在当前包内使用，其他包无法访问。首字母小写的变量在包外部是不可见的。

示例：

```go
package main  

import "fmt"  // 首字母小写的变量只能在当前包内使用 
var privateVariable string = "World"  
func main() {
fmt.Println(privateVariable) // 可以在同一个包内使用     
// fmt.Println(PublicVariable) // 错误，无法访问其他包中的导出变量
}
```

需要注意的是，即使变量名的首字母大写，也并不意味着该变量是公共的，它只是在包之间可见。如果想要将变量作为公共API的一部分，还需要考虑其他因素，如是否提供相应的 getter 和 setter 方法等。

总结起来，首字母大写的变量可以被其他包访问和使用，首字母小写的变量只能在当前包内使用。这种命名规范可以帮助我们更好地组织和管理代码，提高代码的可维护性和可读性。


# 4. Go语言中的异常处理
我熟悉的是Java中的throw new \*\* 以及throws抛出异常以及try...catch...finally操作处理异常操作，对于Go语言中的异常处理有些陌生。

## 1. Go语言中采用defer和panic进行异常操作
`defer` : 用于在函数执行完毕前执行一些清理操作
`panic` : 用于引发异常

当遇到错误或异常情况时，可以使用`panic`函数来引发一个异常。`panic`会停止当前函数的执行，并开始向上返回，**执行每个被延迟的`defer`语句**。如果没有被捕获，异常会导致程序崩溃并打印出调用栈信息。

为了捕获并处理异常，可以使用`recover`函数。`recover`函数用于在`defer`语句中捕获异常，并返回传递给`panic`函数的值。如果没有发生异常，`recover`函数会返回`nil`。

```go
func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("捕获到异常:", err)
        }
    }()

    fmt.Println("开始执行")
    doSomething()
    fmt.Println("执行完成")
}

func doSomething() {
    fmt.Println("执行操作")
    panic("发生异常")
}

```

**异常通常被用于处理无法恢复的错误，而不是用于正常的业务流程控制**

处理流程控制最好可以在发现异常或者错误时返回一个错误码，由错误码来决定程序流程控制，而不是直接利用panic来处理，这样更友好，更易于理解和维护。



## 2. os.Exit()和panic()操作的区别
`os.Exit()`：这个函数会立即终止程序的执行，并返回一个指定的退出状态码。它不会执行任何延迟函数defer()（deferred functions）。  
`panic()` : 这个函数用于表示程序发生了一个无法恢复的错误，会导致程序崩溃。它会触发一个异常，并沿着调用栈向上传递，直到被捕获或者程序终止。在异常被抛出时，会执行所有的延迟函数（deferred functions）。





# 5. gorm操作中会出现数据库名自动后缀加s的操作
比如：struct名字为user，在sql日志中会变成users
此时就需要重写TableName来指定dbname
```go
func (User) TableName() string {  
   //实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）  
   return "user"  
}
```



# 6. gorm的逻辑删除如何实现
## 软删除
`Java`中的`@TableLogic`逻辑删除注解可以理解为Go中的软删除(官方doc是这么叫的)
实现软删除有几种方式：

1. 不嵌套gorm.Model
```go
type User struct {  
	ID      int  
	Deleted gorm.DeletedAt 
	Name    string  
}
```
查找被软删除的记录：
```sql
db.Unscoped().Where("age = 20").Find(&users)  
// SELECT * FROM users WHERE age = 20;
```

2. 永久删除
```sql
db.Unscoped().Delete(&order) 
// DELETE FROM orders WHERE id=10;
```

3. 默认数据类型为时间-> \*time.Time
如果想要用其他数据格式来分别逻辑删除的话（比如0 - 未删除, 1 - 已删除）
通过`import "gorm.io/plugin/soft_delete"`引入插件

```go
  DeletedAt soft_delete.DeletedAt // 利用unix时间戳作为删除标志
  // 利用毫秒milli或者纳秒nano来作为值
  DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli"`  
  DeletedAt soft_delete.DeletedAt `gorm:"softDelete:nano"`
  // 采用 0 和 1作为删除标志（我比较习惯）
  IsDel soft_delete.DeletedAt `gorm:"softDelete:flag"`
```

# 7. Go中的时间格式化问题

直接输出的time.Now()是有显示不友好的问题的
此时需要通过操作：
`time.Now().Format("2006-01-02 15:04:05")`
这样输出的时间才是正常格式化时间

# 8. Go中的类型断言语法
在Go语言中，类型断言用于检查一个接口值的实际类型是否与我们期望的类型匹配
`value, ok := expression.(Type)`

其中，`expression`是一个接口值，`Type`是我们期望的类型。这个语法可以分为两个部分：

1. `expression.(Type)`：这是类型断言的语法。它会尝试将`expression`转换为`Type`类型的值。如果成功，返回转换后的值和`true`；如果失败，返回`Type`类型的零值和`false`。
    
2. `value, ok :=`：这是一个接收返回值的语法。`value`是转换后的值，`ok`是一个布尔值，表示转换是否成功。
```go
func printLength(s interface{}) {
	if str, ok := s.(string); ok {
		fmt.Println("Length of string:", len(str))
	} else if nums, ok := s.([]int); ok {
		fmt.Println("Length of slice:", len(nums))
	} else {
		fmt.Println("Unknown type")
	}
}

func main() {
	printLength("hello")
	printLength([]int{1, 2, 3})
	printLength(42)
}

```

在上面的示例中，`printLength`函数接受一个空接口类型的参数`s`。通过使用类型断言，我们可以检查`s`的实际类型并执行相应的操作。如果`s`是一个字符串类型，我们计算字符串的长度并打印；如果`s`是一个整数切片类型，我们计算切片的长度并打印；否则，我们打印一个未知类型的消息。

# 9. Go如何处理跨域

```go
func Cors() gin.HandlerFunc {  
   return func(c *gin.Context) {  
      method := c.Request.Method  
      origin := c.Request.Header.Get("Origin")  
      if origin != "" {  
         c.Header("Access-Control-Allow-Origin", "http://localhost:8000") // 可将将 * 替换为指定的域名  
         c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")  
         c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")  
         c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")  
         c.Header("Access-Control-Allow-Credentials", "true")  
      }  
      if method == "OPTIONS" {  
         c.AbortWithStatus(http.StatusNoContent)  
      }  
      c.Next()  
   }  
}
```