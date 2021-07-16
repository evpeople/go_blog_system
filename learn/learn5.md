# 不完整的并发

## 基础

1. 除了从主函数退出和直接终止程序以外，没有任何其他的编程方法能够让一个goroutine打断另一个的执行

    1. 可以通过channel，让被请求的goroutine自行结束

    2. > go的时间格式化模板 1 月2 日 3点 4分 5 秒 零六年 UTC-0700

    3. **暂且不能理解的**

        1. go 后面的函数参数在go语句自身被执行的时候求值。

2. channel

    1. 每个channel都有一个类型，表明自己可发送什么类型的数据
        1. ch:= make(chan int)       一个无缓存channel
        2. ch:= make(chan int 5)    一个有缓存channel
        3. ch := make(chan int 0)   一个无缓存channel
    2. chan 传递的也是引用
    3. 运算符是 <-
    4. 通过close(ch) 关闭，关闭后可以接受，不能发送
    5. 无缓存的channel
        1. 发送会导致发送者的goroutine阻塞，直到被接受
        2. 接受会导致接收者的goroutine阻塞，直到被发送
        3. （导致两个goroutine 做一次同步操作，
        4. 当无缓存发送时，接收者收到数据发生在唤醒发送者goroutine前，（**在golang中式保证在唤醒前已经收到了数据，收到数据这件事已经完成了**）
        5. 串联的channel，
            1. 在接受操作时，多接受一个ok值，表明是否channel已经关闭了
            2. 更好的方式时 forr 语句， for x:= range naturals { squares<- x*x  }这样做 naturals是一个channel，完成for后，close(squares)
            3. 只有需要告诉接收者 所有数据已经被发送的时候，才用调用channel。其他的channel会被垃圾回收走
        6. 一般的chan 参数是func xxx(inout chan int)
            1. func xxx( out chan<-  int ) 把int送到chan中（用于发送
            2. func xxx( in <-chan int ) 从chan接受变量到in）（用于接收
            3. 任何双向channel 向单向channel赋值，都会导致一个类型转换
    6. 有缓存的channel（队列的数据结构
        1. 若缓存队列满，则阻塞发送，直到被接受，释放出新的空间
        2. 若缓存队列空，则阻塞接受，直到被填入，获得了新的元素
        3. 可用cap函数获取内部容量，len函数获取有效元素个数
        4. **不能将带缓冲的channel作为同一个goroutine的队列用，可能一个发送操作就永久阻塞了**
        5. **泄露（阻塞在没人接受而卡住的）的goroutine不能被回收**

并发暂且先停止