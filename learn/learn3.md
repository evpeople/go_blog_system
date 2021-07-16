## Panic 异常

1. 可来自运行时的错误，发生会立刻完成defer，然后输出日志。通常不用再次运行程序去定位问题

2. 当到了不可能发生的场景时，采用panic 函数引发异常（断言）

3. 可预料的错误（错误的输入和配置，失败的io，应该人工处理掉，panic尽量避免使用

4. > 函数名中的Must代表不能接受不合法的输入

5. 可用runtime.Stack包输出堆栈信息

## Recover捕获异常

1. 在defer函数中调用了内置的recover，则会使程序从panic中回复，返回panic value，导致panic的函数不会继续运行，但能正常返回，没发生panic时调用recover，则返回nil

    1. panic value时recover返回的

2. 一个例子

3. ```go
   defer func(){
   	if p:=recover();p!=nil{
   		err=fmt.Errorf("internal erroe %v",p)
   	}
       //还可以在其中加入runtime.Stack的信息，添加更多错误信息
   }
   ```



4. 一个规范

    1. 不处理其他包产生的panic，共有的API将函数运行失败当作error返回，而不是panic
    2. 不能不加选择的回复，可以通过不同的panic的返回值来判断是否应该恢复
        1. 也就是说，检测完recover的返回值后，对于完全不知道的panic value再次调用panic（recover的返回值），引发panic。