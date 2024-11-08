# Go learn

## 一、basic 基本语法
### 数据类型
1. bool
2. string
3. byte
4. int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64
5. float32, float64, complex, complex64, complex128
6. rune
7. uintptr 无符号整型没，用于存放一个指针，该类型用于指针计算
8. 结构类型 struct
9. 指针类型
10. 数组
11. 切片 slice
12. map
13. interface{}
14. 通道 channel
15. 函数类型 func
16. 时间类型

### 引用类型
1. 切片 slice
2. map
3. channel
4. interface
5. func
6. 指针类型

## 二、泛型
### golang的泛型使用
1. 泛型函数
2. 泛型类型
3. 泛型接口
4. 泛型结构体
5. 泛型receiver

### 泛型限制
1. 匿名结构体与匿名函数不支持泛型
2. 不支持类型断言
3. 不支持泛型方法，只能通过receiver来实现方法的泛型处理
4. ～后的类型必须为基本类型，不能为接口类型

## 三、关于 `defer、recover、panic`
### 应用场景
1. 资源释放
2. 异常捕获和处理

#### defer
1. defer 关键词用来声明一个延迟调用函数，该函数可以是匿名函数也可以是具名函数
2. defer 延迟函数执行时间（位置），方法 return 之后，返回参数到调用方法之前
3. defer 延迟函数可以在方法返回之后改变函数的返回值
4. 在方法结束（正常返回、异常结束）都会调用defer声明的延迟函数，可以有效避免因异常导致的资源无法释放的问题
5. 可以指定多个defer延迟函数，多个延迟函数的执行顺序为后进先出
6. defer 通常用于资源释放、异常捕获等场景，例如：关闭连接，关闭文件等
7. defer 与 recover 配合可以实现异常捕获与处理逻辑
8. 不建议在for循环中使用defer

#### recover
1. Go语言的内建函数，可以让进入宕机流程中 goroutine 恢复过来
2. recover 仅在延迟函数 defer 中有效，在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果
3. 如果当前的 goroutine 出现 panic，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行

#### panic
1. Go语言的一种异常机制
2. 可通过 panic 函数主动抛出异常

## interface 隐式实现
1. golang 对象实现 interface 无需任何关键词 只需要该对象的方法集中包含接口定义的所有方法且方法签名一致
2. 对象实现接口可以借助 struct 内嵌的特性，实现接口的默认实现
3. 类型T方法集中包含全部 receiver T 方法：类型 *T 方法集包含 receiver T + *T方法
4. 类型T实例 value和pointer可以调用全部的方法，编译器会自动转换
5. 类型T实现接口，不管是T还是*T都实现了该接口
6. 类型 *T 实现接口 只有T类型的指针实现了该接口

## channel 
### 应用场景
1. 协程间通信，即协程间数据传递
2. 并发场景下利用channel的阻塞机制，作为同步机制（类似队列）
3. 利用channel关闭时发送广播的特性，作为协程推出通知

### 通过通讯共享内存
1. channel 的方向，读、写、读写
2. channel 协程间通信信道
3. channel 阻塞协程
4. channel 并发场景下的同步机制
5. channel 通知协程退出
6. channel 的多路复用

> 1. channel 用于协程间通讯，必须存在读写双方，否则会造成死锁

## File 文件读写
### 应用场景
1. 上传下载文件
2. 大文件分片传输
3. 文件移动
4. 文件内容按行获取

### 文件读写
1. 文件复制
2. 一次性读取文件内容并写入新文件
3. 分片读取文件内容分步写入新文件
4. 文件按行读取

## sync.WaitGroup、sync.Cond
### sync.WaitGroup 作用
1. 等待一组协程完成
2. 工作原理：通过计数器来获取协程的完成情况
3. 启动一个协程时计数器+1，协程退出时计数器-1
4. 通过wait方法阻塞主协程，等待计数器清零后才能继续执行后续操作

### sync.WaitGroup 应用场景
1. 通过协程并执行一组任务，且任务全部完成后才能进行下一步操作的情况
2. 如：汽车的生成，所有零件可以并行生产，只能等所有零件生成完成后，才能组装

### sync.WaitGroup 陷阱
1. 协程间传递时需要以指针的方式或闭包的方式引用 WaitGroup 对象，否则将会造成死锁

### sync.Cond 作用
1. 设置一组协程根据条件阻塞，可以根据不同的条件阻塞
2. 可以根据条件唤醒对应的协程

### sync.Cond 应用场景
1. 应用于一发多收的场景，即一组协程需要等待某一个协程完成一些前置准备的情况

### sync.Cond 注意事项
1. 被叫方必须持有锁
2. 主叫方可以持有锁，单允许不持有
3. 尽可能的减少无效唤醒