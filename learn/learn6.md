# golang相关工具

## 包

1. main包是给go build一个信息，此包编译完后必须生产一个可执行文件

2. _test.go 结尾的源文件，且源文件的包名也是以\_test 为后缀的，（前面必须有其他字符

    1. 这些包由go test单独编译

3. 导入包的重名冲突

   ```go
   import{
   	"crypto/rand"
   	mrand "math/rand" // 给这个包重起了给名字，只影响当前的文件
   }
   ```

4. 只想利用导入包的副作用（init初始化函数，计算包级变量的初始化表达式），

    1. 用 _ 重命名导入到包
    2. 用处的例子
        1. 匿名导入 png 包，注册png的解码器，运行的时候能解码这个图像
    3. 数据库也一样

5. 包的命名一般是单数的，避免采用局部变量可能用到的名字

    1. 一个规律是，描述了单一的数据类型的包，一般只暴露一个最主要的数据结构，和相关的方法，以及一个New命名的函数，用于创建一个实例

6. go doc 用于查看文档

    1. 类似javadoc的作用（对于需要大量文档的，建议每个包一个doc.go

    2. go doc time.Since 则显示 Since这个函数的文档，包文档，方法文档同理

       ```go
       // Fprint formats using the default formats for its operands and writes to w.
       // Spaces are added between operands when neither is a string.
       // It returns the number of bytes written and any write error encountered.
       func Fprint(w io.Writer, a ...interface{}) (n int, err error) 
       
       // Package sort provides primitives for sorting slices and user-defined
       // collections.
       package sort
       
       //以被注释的内容作为开头，被注释的内容与注释对象之间没有空行
       // BUG(r): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
       //已知bug,应该反馈给r
       //Deprecated:" 开头的段落 表明不赞成使用，不得不留下
       
       
       ```

7. GOPATH环境变量，用于指定当前的工作目录，不是必须始终如一

8. go instal 与go build

    1. go build 不能生成包文件, go install 可以生成包文件
    2. go build 生成可执行文件在当前目录下， go install 生成可执行文件在bin目录下

9. 内部包

    1. internal包，一个internal包只能被和它有相同父目录的包导入
    2. 包含internal 作为路径的一部分的包就是内部包



## 测试

可通过 -run 配正则式来测试指定函数

1. \_test.go 文件中，有三种函数

    1. 测试函数

        1. Test为函数名前缀的函数，go test调用这些函数，报告结果

        2. 必须导入testing包，签名如下

           ```go
           func TestName(t* testing T){
               //Test 后的后缀必须大写，t参数用于报告测试失败和附加日志信息
               var tests = []struct{
                   input string
                   want bool
               }{
                   {"",true},
                   {"a",true},
                   {"aba",true},
               }
               for _,test:=range tests{
                   if got:=IsPAn(test.input);got!=test.want{
                       t.Errorf("sadsadsa")
                   }
               }
           }
           ```

        3. 一个测试函数往往是调用包内被测试的一个函数，给他一些参数，然后根据返回值判断，若测试出错

           ` t.Error(反引号 xxxxxx（函数名+参数） =false  反引号)`

            1. 也可采用Errrorf 等格式化字符串

        4. 表格驱动的测试（见代码

        5. 一般Error的信息是 *f(x)=y,want z* 若z没有额外信息（比如是布尔量，则可以忽略

    2. 基准测试

        1. 以Benchmark为函数名前缀的测试函数，衡量性能，go test多次运行基准函数计算一个平均时间

        2. 运行时，加上-bench标志

        3. 比较型基准测试样例

           ```go
           func benchmark(b *testing.B size int){}
           func Benchmark10(b *testing.B){ benchmark(b,10) }
           func Benchmark100(b *testing.B){ benchmark(b,100) }
           func Benchmark1000(b *testing.B){ benchmark(b,1000) }
           func Benchmark10000(b *testing.B){ benchmark(b,10000) }
           ```

        4. 运行go tool pprof 刨析一个程序中最该优化的部分

    3. 示例函数

        1. Example为函数名前缀的函数，提供一个由编译器保证正确性的示例文档

        2. Example 的后缀名    会关联到一个具体的（同名）函数上

        3. 如果示例函数内部含有 //Output: 格式的注释，则测试工具会执行示例函数，然后检查标准输出和注释是否一样（应该包括格式？