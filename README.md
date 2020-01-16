# Go语言语法基础
## Label
* [goto](./code/label/goto/main.go)
    - 和C语言一样，可直接跳转到指定标签位置，往下执行
    - 不同于break和continue，goto的标签可以指定为任何语句，包括空语句。
* [break](./code/label/break/main.go)
    - 不同于goto标签，break标签只能指定`for`、`switch`和`select`语句，并且`break`语句只能跳转到当前语句的外部，如下面的代码时无法编译通过的：
    ```
    FirstLoop:
        for i := 0; i < 10; i++ {
        }
        for i := 0; i < 10; i++ {
            break FirstLoop
        }
    ```
* [continue](./code/label/continue/main.go)
    - 不同于break标签，continue标签只能修饰`foo`，因为另外两个没办法循环执行

# Go语言各种包的用法
## flag
* [flag](./code/flag/main.go)
    - 定义flag名字、默认值、描述已经内容接收变量 -> `flag.Parse()` -> 操作接收呢容
    - 命令行数据格式：
    > go run main.go -flagname=flagvalue

## context
* [context-for-sync](./code/context/sync.go)
    - context可以用做主线程控制子线程的通信机制
    - 有三种context：cancelContext, timeoutContext, deadlineContext
    - cancelContext需要主动调用`cancel`，信号量才会`Done`。另外两种除了主动调用`cancel`外，时间到了，也会`Done`。
* [context-for-data](./code/context/value.go)
    - `context.WithValue()`是向子线程传递数据的方法，由于没有`cacel`机制，无法控制子线程的运作。

# Gin Web Framework
## 接收用户传过来的数据
* [gin-input-from-url-query-string](./code/gin/input/query-string.go)
    - 在浏览器中输入:
    > localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15
    - URL可以带[query string](https://en.wikipedia.org/wiki/Query_string)的
    - 对GET请求，gin只启用`Form`组件解析`query string`
    - `c *gin.Context`是回调函数接收的参数
    - `c.ShouldBind(&person)`接受`gin.Context`赋值
    - `type Persion struct`定义了`Form`结构，定义key-value，和顺序无关
## 发送JSON数据给客户
* [gin-output-with-json](./code/gin/output/main.go)
    - 在浏览器中输入以下测试数据：
    ```
    1.http://localhost:8080/getb?field_a=hello&field_b=world
      输出：{"a":{"FieldA":"hello"},"b":"world"}
    2.http://localhost:8080/getc?field_a=hello&field_c=world
      输出：{"a":{"FieldA":"hello"},"c":"world"}
    3.http://localhost:8080/getd?field_x=hello&field_d=world
      输出：{"d":"world","x":{"FieldX":"hello"}}
    ```
    - 第一层的key，代码指定
    - 第二层的key，代码根据结构体定义推到出来
    - 嵌入的JSON格式，由三种指定方法：1.内部用结构体；1.内部用结构体指针；2.内部定义结构体；
## 中间件
* [gin-middleware-process](./code/gin/middleware/main.go)
    - 新建一个空的gin引擎，`gin.Default()`默认回加上gin自己的Logger
    - 添加中间件`Logger()`，中间件是一个`HandlerFunc`
    - 在中间件中，`c.Next()`之前都是Router回调函数之前执行的，之后都是Router回调函数之后执行的
    - `c.Set()`可以往context中填入[key, value]

# Logger logrus
## 基本输出格式
* [logger-logrus-hello](./code/logrus/hello/main.go)
    - 基本输出格式为（field:[animal, walrus]）：
    > time="2020-01-16T19:16:03+08:00" level=info msg="A walrus appears" animal=walrus
* [logger-logrus-json](./code/logrus/customize/json-level.go)
    - 可设置输出格式和输出级别
* [logger-location-file](./code/logrus/location/file.go)
    - 可创建多个logger
    - 可方便输出到文件
## Hooks
* `logrus`提供了日志钩子，可以挂上数据库日志操作的钩子函数，这样`logrus`不仅本地能将log打到本地的窗口，也会打印到数据库中。
* [logrus-log-mongodb](./code/logrus/hook/mongodb.go)
    - `logrus`的`Hook`接口
    ```go
    type Hook interface {
        Levels() []Level
        Fire(*Entry) error
    }
    ```
    - `github.com/LyricTian/logrus-mongo-hook/hook.go`实现了这个接口
    - 因此，`mongohook.DefaultWithURL()`可以产生一个`Hook`，挂到`logrus`上，即可将日志写到数据库中了
    - 启动程序后，mongodb会多出一条文档信息：
    > { "_id" : ObjectId("5e20957ab48172e81ec607a7"), "foo-warn" : "bar-wan", "level" : 3, "message" : "test warn", "created" : NumberLong(1579193722) }
