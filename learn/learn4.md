# golang面向对象

## 方法

1. 方法的声明

    1. 在函数声明时，函数名字前放上一个变量就是一个方法，相当于为这个变量的类型附加了一个方法

    2. 用古早的话，叫此方法的接收器

     ```go
      type Point struct{ X, Y float64 }
      
      // traditional function
      func Distance(p, q Point) float64 {
      	return math.Hypot(q.X-p.X, q.Y-p.Y)
      }
      
      // same thing, but as a method of the Point type
      func (p Point) Distance(q Point) float64 {
      	return math.Hypot(q.X-p.X, q.Y-p.Y) //注意此处的p
      }
      // A Path is a journey connecting the points with straight lines.
      type Path []Point
      
      // Distance returns the distance traveled along the path.
      func (path Path) Distance() float64 {
      	sum := 0.0
      	for i := range path {
      		if i > 0 {
      			sum += path[i-1].Distance(path[i])
      		}
      	}
      	return sum
      }
      ```

    4. 这个接收器相当于其他语言的self or this，我们可以自己选择，惯例是类型的首字母

    5. **注意，一个方法的名字不能和结构体的成员变量的名字重复（因为是在同一命名空间下**

    6. 从上文的path可看出，不只是能给struct类型定义方法，实际上可以给任意（非 指针，非 interface）类型定义方法，

    7. 一个类型的方法名必须是唯一的，不能通过参数列表和返回值区分

    8. 当接受者变量较大的时候，用接受者的指针

    9. 几个约定

        1. **一个类有一个指针作为接收者的方法，则他的所有方法都必须有一个指针接收器，即使是不需要的函数**

        2. **当类型名为指针的话，不允许作接收器（也就是不能给底层是指针的类型定义方法**

        3. **nil作为合法的接收者的时候，应在注释中说明nil的逻辑意义**

            1. 调用的时候应该写` Value(nil).Get`  显式的对nil做类型转换

        4. golang可以在用指针调用方法的时候，不显式的说明自己是指针，例子见下面

        5.

        ```go
        p:=&Point{1,2}
        p.ScaleBy()
        
        //也可使用
        p:=Point{1,2}
        p.ScaleBy()
        
        //这两个调用的其实是一个方法，但是这种简写只适用于变量，临时变量（Point{1,2}.ScaleBy）不可
        ```

        6. > 总结：
           >
           > 1. 不管方法的recviver是指针还是非指针，都可以通过指针或非指针调用，编译器类型转换
           >
           > 2. 声明一个方法的接收者的时候，
                >    1. 要考虑此对象是否很大，会在拷贝的时候产生较大的消耗，
           >    2. 指针的指向的始终是一个内存地址，进行了拷贝可能影响内存不变性（cpp的自己实现深拷贝

    10. 内嵌的类型的方法，相当于组合到了新的结构体里，是一个组合的关系（设计模式中推荐采用组合而不是继承）

        1. 对于匿名内部类的方法的使用是采用包装的设计模式，自动生成了包装好的本类方法。
        2. 同时可手动重写方法
        3. 从首级开始找方法名，然后递归的向下寻找方法，方法名有二义性的时候会报错（两个内部匿名类不能有同名方法，或者采用显式的调用

    11. 内嵌一个指针，一个有趣的应用是两个圆共享一个圆心（内存地址相同的圆心）

    12. 另一个有趣的匿名类的应用

        ```go
          var (
          	mu sync.Mutex
              mapping=make(map[string]int)
          )
          
          func Lookup(key string)int{
              mu.Lock()
              v:=mapping[key]
              mu.Unlock()
              return v
          }
          
          
          var cache=struct{
              sync.Mutex
              mapping map[string]int
          }{
              mapping:make(map[string]int),
          } //匿名结构体要在声明后初始化
          
          func Lookup(key string)int{
              cache.Lock() //魔法
              v:=cache.mapping[key]
              cache.Unlock()
              return v 
          }
          ```

        2. > 再次强调，匿名结构体要在声明后初始化，即便是{}这样的空初始化

    13. 方法值和方法表达式

        1. p.Distance() 在p上调用Distance方法
        2. p.Distance  给Distance选择接受对象为p
            1. 使用方法 disP:=p.Distance
            2. disP(q)
            3. 可以不指定接收者调用（因为已经指定过了
            4. 用处是将函数作为一等公民（高阶函数）传入其他函数的时候，已经确定好了接收者，
                1. 比如，省略一个匿名函数
        3. 一个很好的可以用到的地方是大一上的模拟冯诺依曼计算机，给指令类型添加Add Sub 方法 ，然后op=指令.Add 后面的都用op操作操作数就可以了

    14. > 用bit数组（Slice）也可表示集合

    15. fmt包的函数会调用用户给类型的String方法
## 封装
1. go中控制可见性的手段只有大小写的标识符
2. 最小封装单元是Package而不是类
3. 对于getter 方法，golang倾向于直接忽略掉Get前缀，用被获取的成员为方法命名
    1. 虽然没有禁止直接通过名字获取而不是方法获取，但是方法获取总是更好一些

## 接口

### 接口的基础知识

1. 对接口类型，只能知道可以用他的方法来做些什么

2. 接口就是一个约定，约定被调用者提供什么方法，且保证调用者能够正确使用这个方法(方法的接收者类型并不重要)

3. LSP里氏替换

4. 之所以定义一个String方法，就能够打印这个类型，是因为定义了String方法，使得值满足了Stringer类型

5. 接口类型详解

    1. 具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例

    2. 可通过组合形成新的接口

    3. 一个类型若拥有一个接口所需要的所有方法，则这个类型就实现了这个接口

        1. 可以用接口更抽象的描述的类型，所以go程序员常把一个具体类型描述为一个特定的接口类型

         ```go
         var w io.Writer()
         /*
         	则w可以被赋值为os.Stdout , bytes.Buffer,time.Second 
         	甚至可以被赋值为其他（包含了）了Write方法的接口
         */
         ```

6. 虽然在方法调用的时候，语法糖让人可以不管接收者是不是指针，但是若一个指针接收者实现了某方法而类型本身没有实现，则指针接收者就实现了更大的接口

7. 接口封装了具体类型和值，具体类型也只有接口中包含的方法能够被调用

    1. 比如把os.Stdout 赋值给 var w io.Writer 则w 不能调用Close方法

8. 空接口类型的变量可以被赋任何值interface{}   (通过类型断言获取空接口中的方法)

9. 一个编译期断言 某只实现了某接口的例子

    1.

   ```go
   var _ io.Writer=(*bytes.Buffer)(nil)
   ```

10. 一个接口方法隐式的给接收者带来变化的时候，用指针承载方法

11. go语言可以在需要的时候定义一个接口集合

12. 接口值是可以用比较的，当他们都是nil值或者他们的动态类型相同且动态值也相同的的时候，（可用于map和swtich

    1. 但是动态值不可比时，会出panic （切片
    2. 只有接口是这样的，其他类型要么安全可比，要么完全不可比，就这个比较的时候可能报panic

13. **一个包含nil指针的接口不是nil接口**

    1. nil接口的类型也是nil
    2. 可能产生问题是因为  某个类型的指针可能确实实现了此接口的方法，但是不代表此类型的nil指针也满足这个方法正常使用的条件

### 接口的具体例子

#### sort.Interface接口（伪泛型

1. 此接口中包含了三个方法

   ```go
   package sort
   
   type Interface interface{
   	Len()int
   	Less(i,j int)bool //之所以是int 是代表了slice下标
   	Swap(i,j int)
   }
   ```

2. 定义一个实现了上述三个方法的类型

3. 将一个普通切片类型转化为上述类型

4. 通过sort.Sort(StringSlice(names)) 对names 排序，实际的排序是Sort的函数，这个函数的参数类型是sort.Interface

5. 对一个音乐列表通过多种方式排序的时，应该将type byYear []*Track 定义多种排序的类型，然后每个排序类型分别实现上述接口，排序的时候转化为相应的类型

    1. 也可以用一个结构体，包含一个函数成员，返回实现Len（） Swap（），对于Less 的实现是调用结构体中的函数成员，每次排序的时候，生成此结构体（把待排序的slice作为第一个成员，函数成员用匿名函数

### http.Handler接口

1. 此接口的方法

   ```go
   type Handler interface{
   	ServerHTTP(w ResponseWriter r* Request)
   }
   ```

2. 网络编程

    1. 通过ServerMux简化URL和handlers的联系，一个ServerMux将一批http.Handler聚焦到一个http.Handler中
    2. 先http.NewServeMux()，做出一个mux然后mux.Handle("/list",http.HandlerFunc( 实际函数))   也就是传入一个前缀，传入一个handlerFunc

   ```go
   mux:=http.NewServeMux()
   mux.Handle("\list",http.HandlerFunc(db.list))
   
   func (db datebase)list(w http.ResponseWriter,req *http.Request)
   {
   ....
   }
   
   /*此处的的HandlerFunc 做的是，把list这种参数列表的函数做类型转换，转换为HandlerFunc类型（这种函数），而这种函数类型，实现了Server方法（通过包装） ,从而实现了http.Handler接口，从而可以被调用*/ 
   
   // 实际上handleFuncr是一个适配器
   
   /*用书中的言论来说，db.list是一个函数，但是没有方法，所以没有实现hadler接口，所以不能传给mux.Handler*/
   ```

3. 在实际编程中，常用的是mux.HandleFunc("/list",db.list) 简化一个函数

4. 实际上，直接使用handleFunc，而不使用mux，会将Func添加到默认handler里，此时启动服务器时，其中的handler接口的参数时nil

5. 当handler获取其他协程 （可访问的变量）or本身的其他请求（可访问的变量）的时候，要加锁（或者其他预防措施）

### error接口

1. 一般使用fmt.Errorf 进行error的说明（此函数调用了errors.New

## 类型性质

### 类型断言

1. 类型断言（一般用于扩大可获取的方法集合，当断言的类型是一个接口的时候

    1. 语法 w.(io.Writer)
    2. 若断言成功，则w 可以写了

2. 当断言类型是一个具体类型时，断言成功后，返回原本的值，断言失败时panic

3. 通过将Error类型用不同的结构体包装起来，然后在新的结构体实现Error接口

    1. 接着通过类型断言将不同的失败区分开来

4. > 向io.Writer写入字符串推荐的函数是io,WriteString 这个函数内部声明了一个stringWriter的接口，然后会断言传入的io.Writer也是stringWriter接口，若断言成功则调用WriterString，断言失败再使用w.Write()

###  类型开关

1. 接口使用方式
    1. 表达实现这个接口的类的相似性，但是隐藏具体操作（子类型多态（上文提到的三种类型
    2. 认为接口是类型的联合，然后通过断言动态区别接口（非参数多态
2. 类型开关就是switch x.(type) 的switch case
    1. 当多个case是接口时，case的顺序就很重要，因为他可能满足多个接口
    2. 扩展形式是 switch x:x.type() 将提取的值绑定到一个在switch case范围内的变量



## ***接口使用的建议！***

1. 不应该在设计一个新包的时候，先创建一套接口，然后再定义满足他的具体实现，只有当两个或以上的具体类型必须以相同方式处理时才需要接口

2. 除非具体类型不能和接口存在于同一个包中，这时候可以用一个接口解耦两个包

3. 设计接口的好的标准是ask only for what you need 只考虑要用的

   
