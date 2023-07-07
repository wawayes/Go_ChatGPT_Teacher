1. Init函数

在 Go 中，`func init()` 是一个特殊的函数，它用于在程序启动时自动执行一些初始化操作。`init()` 函数没有参数和返回值。

`init()` 函数在 Go 程序中的每个包中都可以存在，且可以存在多个。当程序启动时，Go 运行时会自动按照包的导入顺序调用每个包中的 `init()` 函数。

`init()` 函数的作用可以包括但不限于：

1. 初始化全局变量或常量。
2. 执行一些需要提前完成的操作，例如注册数据库驱动、初始化日志等。
3. 进行一些需要在程序启动时执行的设置，例如解析命令行参数、读取配置文件等。

需要注意的是，`init()` 函数是在程序启动阶段自动执行的，而不是在创建实例时自动构造。它与结构体的构造函数不同。`init()` 函数在程序启动时只执行一次，而结构体的构造函数是在创建实例时被调用。

2. 变量首字母大写｜小写

3. os.Exit() | panic()


4. gorm指明数据库表名

5. gorm逻辑删除


6. gorm时间格式化
https://studygolang.com/articles/35739

7. golang时间格式化

go语言诞生的时间：
```go
"2006.01.02 15:04:05"
```

8. redis: can't marshal map[string]interface {} (implement encoding.BinaryMarshaler)

9. go web session轮子
https://juejin.cn/post/6952340347280162852

10. go 鉴权方式
- HTTP Basic Authentication
- Session-Cookie机制
- Token令牌机制
- OAuth2.0授权机制

11. 类型断言语法

12. 跨域
