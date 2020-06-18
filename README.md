# Go语言圣经学习

## Go语言程序结构
### 命名
Go语言中的函数名、变量名、常量名、类型名、语句标号和包名等所有的命名，都遵循一个简单的命名规则：一个名字必须以一个字母(Unicode字母)或
下划线开头、后面可以跟任意数量的字母、数字或下划线。Go语言中大写字母和小写字母开头的常量、变量和函数名等是不同的，用以区分公有、私有。

Go语言中类似if和switch的关键字有25个；关键字不能用于自定义名字，只能在特定语法结构中使用。
```
break       default     func        interface       select
case        defer       go          map             struct
chan        else        goto        package         switch
const       fallthrough if          range           type
```

Go语言中包含大约30多个预定义的名字：
```
内建常量：true false iota nil
内建类型：int int8 int16 int32 int64
        uint uint8 uint16 uint32 uint64 uintptr
        float32 float64 complex128 complex64
        bool byte rune string error
内建函数：make len cap new append close copy 
        delete complex real imag panic recover
```
这些内部预先定义的名字不是关键字，可以在定义中重新使用它们。如果一个名字是在函数内部定义，那么它的就只在函数内部有效。如果是在函数外部定义，
且该名字是大写字母开头的，表示是可以被外部的包访问，否则不能；函数大写字母开头和小写字母开头也类似。

### 声明
声明语句定义了程序的各种实体对象以及部分或全部的属性。Go语言主要有四种类型的声明语句：var、const、type和func，分别对应变量、常量、类型和
函数实体对象的声明。

一个Go语言编写的程序对应一个或多个以.go为文件后缀名的源文件中。每个源文件以包的声明语句开始，说明该源文件属于哪个包。包声明语句之后是import
语句导入依赖的包，然后是包一级的类型、变量、常量、函数的声明语句，包一级的各种类型的声明语句的顺序无关紧要。

一个函数声明由一个函数名字、参数列表、一个可选的返回值列表和包含函数定义的函数题组成。如果函数没有返回值，那么返回值列表是可以省略的。执行函数
从函数的第一个语句开始，依次顺序执行直到遇到return返回语句，如果没有返回语句则是执行到函数末尾，然后返回到函数调用者。
