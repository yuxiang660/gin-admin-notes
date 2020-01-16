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
