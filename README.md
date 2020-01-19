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
## struct、指针、接口
* [error-pointer-prog](./code/pointer/main.go)
    - 不同于C++，Go是不支持函数重载的，而指针类型函数和原类型函数如果定义一样，是算一个函数的，会编译报错`method redeclared`
    - 一个接口只关心函数行为，而不关心当前函数是指针函数，还是非指针函数。如果是指针函数，就有可能影响当前指针指向对象的内容。如果不是指针函数，无论这个函数里面做了什么行为，原类型一直保持不变。

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
## 静态文件路由配置
* [gin-static-router-server](./code/gin/static/main.go)
    - 有三种方法可以配置静态服务路由

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

# GORM
## Hello
* [gorm-sqlite-hello](./code/gorm/hello/main.go)
    - 安装好sqlite即可运行此程序
    - gorm只支持关系型数据库
    - 连接数据库 -> 绑定数据模型 -> 操作数据库（CQUD）

# JWT
## JWT概念
* JWT全称JSON Web Token，是一种加密算法，常用于互联网请求的验证
* 实现机制
    - 用户携带用户信息，请求登录服务器
    - 服务器验证用户信息，如果通过，则产生JWT，并通过Cookie发送回客户端
    - 客户端收到服务器的回复，存下Cookie
    - 用户发送其他请求到服务器时，需要携带此Cookie
    - 服务器处理此请求时，先要验证此Cookie是否有效，才能继续执行
    - 验证有效后，执行用户请求，并返回
    ![Token-Based-Auth.png](./code/jwt/Token-Based-Auth.png)
* JWT-Token和Session的区别和联系
    - 参考[link](https://medium.com/@sherryhsu/session-vs-token-based-authentication-11a6c5ac45e4)
    - Session验证的`SessionID`是存在服务器端的，而Token验证由Token自己就可以完成，不需要再利用服务器端的存储信息进行验证。
    - 不需要在服务器端存储验证信息，是JWT-Token方式的最大优势。
* 示例代码[go-jwt-server](./code/jwt/server/main.go)，参考[link](https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/)
    - 启动服务
    - 用`Postman`模拟客户端行为
        - 发送Get请求到：http://localhost:8000/welcome
        > 收到`401 Unauthorized`错误，因为我们没有登录验证
        - 发送Post请求到：http://localhost:8000/login , 并带上JSON格式的用户户名密码`{"username":"user1","password":"password2"}`。服务器写死了两个用户：user1和user2。
        > 验证成功，服务器发送了Cookie给客户端
        - 打开Postman的`Cookies`串口，发现`localhost`下多了一个`token`Cookie，这就是服务器发送过来的，里面由token信息
        - 再次发送Get请求到：http://localhost:8000/welcome
        > 收到 `Welcome user1!`回复
        - 删除Postman的`token`Cookie
        - 再次发送Get请求到：http://localhost:8000/welcome
        > 收到`401 Unauthorized`错误，因为客户端的Cookie被删除了，需要再次登录验证
    - login HTTP POST包，申请验证用户
        - client -> server
        ```
        POST /login HTTP/1.1\r\n
        Content-Type: application/json\r\n
        User-Agent: PostmanRuntime/7.21.0\r\n
        Accept: */*\r\n
        Cache-Control: no-cache\r\n
        Host: 192.168.50.64:8000\r\n
        Accept-Encoding: gzip, deflate\r\n
        Content-Length: 43\r\n
        Connection: keep-alive\r\n

        {"username":"user1","password":"password1"}
        ```
        - server -> client
        ```
        HTTP/1.1 200 OK\r\n
        Response Version: HTTP/1.1
        Status Code: 200
        Response Phrase: OK
        Set-Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTc5MjUzODI2fQ.NYYRmICsUPURGXndOt7gjGvX-x3tvZsjTEyjR453g_M; Expires=Fri, 17 Jan 2020 09:37:06 GMT\r\n
        ```
    - 收到验证Cookie后，客户端向服务器发送welcome HTTP GET包
        - client -> server
        ```
        GET /welcome HTTP/1.1\r\n
        User-Agent: PostmanRuntime/7.21.0\r\n
        Accept: */*\r\n
        Cache-Control: no-cache\r\n
        Postman-Token: b1f6a751-7041-4abf-895b-c0efb56b1dd9\r\n
        Host: 192.168.50.64:8000\r\n
        Accept-Encoding: gzip, deflate\r\n
        Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTc5MjUzODI2fQ.NYYRmICsUPURGXndOt7gjGvX-x3tvZsjTEyjR453g_M\r\n
        Connection: keep-alive\r\n
        ```
        - server -> client
        ```
        HTTP/1.1 200 OK\r\n
        Content-Length: 14\r\n
        Content-Type: text/plain; charset=utf-8\r\n

        Welcome user1!
        ```

# OAuth2
## OAuth2概念
OAuth是一种授权机制。数据的所有者告诉系统，同意授权第三方应用进入系统，获取这些数据。系统从而产生一个短期的进入令牌（token），用来代替密码，供第三方应用使用。参见[Link](http://www.ruanyifeng.com/blog/2019/04/oauth_design.html)
## OAuth2和JWT的关系
个人理解，JWT是一种Token生成和校验的方法，OAuth2也用到了技术。OAuth2利用JWT技术，实现了第三方授权机制的验证模式。OAuth2的实现机制有多种选择。
* 上面JWT实现的保护和OAuth2的相同点
都生成Token给客户端。目的是，提供一个短期令牌（而不是密码）给客户端，让客户端可以在短期内利用这个令牌，不再用用户授权，就可以访问受保护的资源。
* 上面JWT实现的保护和OAuth2的不同的地方
JWT方法需要客户端提供用户名和密码（或者说登录后），才提供Token，相当于把用户和客户端看成一个人了。而OAuth将客户端和用户分离，客户端和用户之间通过权限服务器沟通，用户授权权限服务器可以给此客户端Token，服务器就生成Token给此客户端用。这样做的好处是，用户的用户名和密码只有用户知道，权限服务器和客户端都不知道。
## 示例代码
* [go-oauth-simple-demo](./code/oauth/main.go)
* 此代码是简化的示例，用户授权的过程体现在事先定死的`CLIENT_ID`和`CLIENT_SECRET`。只要权限服务收到此客户端的Token申请需求，就生成Token给客户端。
* 执行流程：
    - 开启服务
    > go run main
    - 浏览器申请Token
    > localhost:9094/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=all
    - 浏览器收到回复，例如
    > {"access_token":"T2LMWRAWO_M3MBLG17S-UQ","expires_in":7200,"scope":"all","token_type":"Bearer"}
    - 利用上面的Token，操作资源
    > http://localhost:9094/protected?access_token=T2LMWRAWO_M3MBLG17S-UQ
    - 收到回复
    > Hello, I'm protected

# Dig 模块
* [go-dependency-inject-dig](./code/dig/main.go)，参考[link](https://blog.drewolson.org/dependency-injection-in-go)
    - Dig的适用范围：当某对象的依赖较多且逻辑复杂时，可将依赖的构造和初始化的工作交给Dig，通过依赖注入的方式，构造此对象。
    - Dig实现流程：
        - 注册依赖构造函数：将所有依赖的构造函数注册给Dig(通过Provide)，单个依赖的构造可以依赖其他依赖。注意，这一步并不会执行依赖的构造（真正的构造是在第一次调用`invoke`的时候），而且Provide调用的顺序也不影响后面的依赖构造顺序。
        - 完成依赖注册：Dig返回一个`Container`，包含所有依赖的信息。此时，所有的依赖依旧还没有构造。
        - 调用依赖函数：当第一次调用`Container.Inoke(...)`时，构造所有依赖。`Inoke`函数用于调用所有依赖的方法。

# Redis和Buntdb数据存储模块
* [go-redis](./code/redis/main.go) 和 [go-buntdb](./code/buntdb/main.go)
    - 两者都支持`key-value`的存储
    - `buntdb`支持直接存入文件，而且不需要额外安装`redis`软件
    - `redis`对大量数据效率可能更好

