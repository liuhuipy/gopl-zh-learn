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
continue    for         import      return          var
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

### 变量
var关键字声明语句可以创建一个特定类型的变量，然后给变量附加一个名字，并设置初始值：
```
var 变量名字 变量类型 = 初始化表达式
var a int = 10
var b string = "hello"
var c names []string 
```
其中"变量类型"和"= 初始化表达式"可以省略其中一个。如果省略类型信息，那么将根据初始化表达式来推导变量的类型信息。如果初始化表达式被省略，那么
将用变量类型的零值来初始化该变量。数值类型变量的零值是0，字符串类型的零值是""，接口或引用类型（包括slice、指针、map、chan和函数）对应
的零值是nil，布尔类型的零值是false。一组同类型的变量可以简化为只在最后一个变量后面声明类型。
```
var i, j, k int                     // 初始化值 0, 0, 0
var m, n string                     // 初始化值 "", ""
var b, f, s = true, 2.3, "four"     // 对应类型 bool, float64, string
```
简短变量声明：
```
i := 100
a, b, c := 2.2, "world", true
```
#### 指针
一个变量对应一个保存了变量对应类型值的内存空间。普通变量在声明语句创建时被绑定到一个变量名，比如叫x的变量，但是还有很多变量始终以表达式方式
引入，例如x[i]或x.f变量。所有这些表达式一般都是读取一个变量的值，除非它们是出现在赋值语句的左边，这种时候是给对应变量赋予一个新的值。

一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。并不是每一个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址。
通过指针，我们可以直接读或更新对应变量的值，而不需要知道该变量的名字。

如果用"var x int"声明一个x变量，那么&x表达式（即取x变量的内存地址）将产生一个指向该整数变量的指针，指针对应的数据类型是*int，指针被称之为
"指向int类型的指针"。如果指针名字为p，那么可以说"p指针指向变量x"，或者说"p指针保存了x变量的内存地址"。同时*p表达式对应p指针指向的变量的值。
一般*p表达式读取指针指向的变量的值，这里为int类型的值，同时因为*p对应一个变量，所以该表达式也可以出现在赋值语句的左边，表示更新指针所指向的
变量的值。
```
x := 1
p := &x             // p type *int
fmt.Println(*p)     // 1
*p = 2
fmt.Println(x)      // 2
```
对于聚合类型每个成员--比如结构体中的每个字段、或者是数组的每个元素--也都是对应一个变量，因此可以被取地址。变量有时候被称为可寻址的值。即使
变量由表达式临时生成，那么表达式也必须能接受&取地址操作。任何类型的指针的零值都是nil，如果p指向某个有效变量，那么p != nil测试为真。指针
之间也是可以进行相等测试的，只有当它们指向同一个变量或全部是nil时才相等。
```
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil)  // true false false
```
在Go语言中，返回函数中局部变量的地址也是安全的。例如下面的代码，调用f函数时创建局部变量v，在局部变量地址被返回之后依然有效，因为指针p依然
引用了这个变量。
```
var p = f()

func f() *int {
    v := 1
    return &v
}

# 每次调用f函数都将返回不同的结果
fmt.Println(f() == f())         // false
```
因为指针包含了一个变量的地址，因此如果将指针作为参数调用函数，那将可以在函数中通过该指针来更新变量的值。例如下面这个例子就是通过指针来更新
变量的值，然后返回更新后的值，可用在一个表达式中
```
func incr(p *int) int {
    *p++
    return *p
}

v := 1
incr(&v)                // 2
fmt.Println(incr(&v))   // 3
```
每次我们对一个变量取地址，或者复制指针，我们都是为原变量创建了一个新的别名。例如，*p就是变量v的别名。指针特别有价值的地方在于我们可以不用名字
而访问一个变量，但是这是一把双刃剑，要找到一个变量的所有访问者并不容易，我们必须知道变量全部的别名。不仅仅是指针会创建别名，很多其它引用类型
也会创建别名，例如：slice、map和chan，甚至结构体、数组和接口都会创建所引用变量的别名。

指针是实现标准库中flag包的关键技术，它使用命令行参数来设置对应变量的值，而这些对应命令行标志参数的变量可能会零散的分布在整个程序中。为了说明
这一点，在早些的echo版本中，就包含了两个可选的命令行参数：-n用于忽略行尾的换行符，-s sep用于指定分隔字符（默认为空格）。下面这是第四个版本，
对应包路径为gopl-zh-learn/ch2/echo4：
```go
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(*n, *sep)
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
```
调用flag.Bool函数会创建一个新的对应布尔型标志参数的变量。它有三个属性：第一个是命令行标志参数的名字"n"，第二个是该标志参数的默认值，第三个
是该标志参数对应的描述信息。如果用户在命令行输入了无效的标志参数，或者输入-h或-help参数，那么将打印所有标志参数的名字、默认值和描述信息。类似
的，调用flag.String函数将创建一个对应字符串类型的标志参数变量，同样包含命令行标志参数对应的参数名、默认值和描述信息。程序中的sep和n变量分别
是指向对应命令行标志参数变量的指针，因此必须用*sep和*n形式的指针语法间接引用它们。

当程序运行时，必须在使用标志参数对应的变量之前先调用flag.Parse函数，用于更新每个标志参数对应变量的值（之前是默认值）。对于非标志参数的普通
命令行参数可以通过调用flag.Args()函数来访问，返回值对应一个字符串类型的slice，如果在flag.Parse函数解析命令行参数时遇到错误，默认将打印
相关的提示信息。

#### new函数
另一个创建变量的方法是调用内建的new函数。表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T。
```
p := new(int)
fmt.Println(*p)         // 0
*p = 2
fmt.Println(*p)         // 2
```
用new创建变量和普通变量声明语句方式创建变量没什么区别，除了不需要声明一个临时变量的名字外，我们还可以在表达式中使用new(T)。换而言之，new函数
是一种语法糖，而不是一个新的基础概念。下面的两个newInt函数有着相同的行为：
```
func newInt() *int {
    return new(int)
}

func newInt() *int {
    var dummy int
    return &dummy
}
```

#### 变量的生命周期
变量的生命周期指的是在程序运行期间变量有效存在的时间间隔。对于在包一级声明的变量来说，它们的生命周期和整个程序的运行周期是一致的。而相比之下，
局部变量的生命周期则是动态的：每次从创建一个新变量的声明语句开始，直到该变量不再被引用为止，然后变量的存储空间可能被回收。函数的参数变量和返
回值变量都是局部变量。它们在函数每次被调用的时候创建。
```
for t := 0.0; t < cycles*2*math.Pi; t += res {
    x := math.Sin(t)
    y := math.Sin(t*freq + phase)
    img.SetColorIndex(
        size+int(x*size+0.5), size+int(y*size+0.5),
        blackIndex,
    )
}
```
在每次循环的开始会创建临时变量t，然后在每次循环迭代中创建临时变量x和y。
那么Go语言的自动垃圾收集器是如何直到一个变量是何时可以被回收呢？这里我们可以避开完整的技术细节，基本的实现思路是，从每个包级的变量和每个当前运行
函数的每一个局部变量开始，通过指针或引用的访问路径遍历，是否可以找到该变量。如果不存在这样的访问路径，那么说明该变量是不可达的，也就是说它是否
存在并不会影响程序后续的计算结果。因为一个变量的有效周期只取决于是否可达，因此一个循环迭代内部的局部变量的生命周期可能超出其局部作用域。同时，局部
变量可能在函数返回之后依然存在。编译器会自动选择在栈上还是在堆上分配局部变量的存储空间，这个选择并不是由用var还是new声明变量的方式决定的。
```
var global *int

func f() {
    var x int
    x = 1
    gloabl = &x
}

func g() {
    y := new(int)
    *y = 1
}
```
f函数里的x变量必须在堆上分配，因为它在函数退出后依然可以通过包一级的global变量找到，虽然它是在函数内部定义的；用Go语言的术语来说，这个x局部变量
从函数中逃逸了。相反，当g函数返回时，变量*y将是不可达的，也就是说可以马上被回收的。因此，*y并没有从函数g中逃逸，编译器可以选择在栈上分配*y的存储
空间，也可以选择在堆上分配，然后由Go语言的GC回收这个变量的内存空间。虽然这里用的是new方式，其实在任何时候，你并不需为了编写正确的代码而要考虑变量
的逃逸行为，要记住的是，逃逸的变量需要额外分配内存，同时对性能的优化可能会产生细微的影响。

### 赋值
使用赋值语句更新一个变量的值：
```
x = 1                           // 命名变量的赋值
*p = true                       // 通过指针间接赋值
person.name = "bob"             // 结构体字段赋值
count[x] *= scale     // 数组、slice或map的元素赋值

v := 1
v++
v--                 // 数值变量支持++递增和--递减语句

i, j, k = 2, 3, 5
x, y = y, x
a[i], a[j] = a[j], a[i]
```

### 自定义类型
用一个类型声明语句可以创建一个新的类型名称，和现有类型具有相同的底层结构：
```
type 类型名称 底层类型

type Celsius float64
type Fahrenheit float64
```
两个不同类型的值不可以进行比较，即使它们底层类型一样：
```
var c Celsius
var f Fahrenheit
fmt.Println(c == 0)         // true，Celsius和Fahrenheit类型的底层类型都是float64,零值为0
fmt.Println(f == 0)         // true
fmt.Println(c == f)         // compile error: type mismatch
fmt.Println(c == Celsius(f))    // true，转换为统一类型后可以进行比较
```

### todo: 包和文件
### todo: 作用域

## 基础数据结构
从底层而言，所有的数据都是由比特组成，但计算机一般操作的是固定大小的数，如整数、浮点数、比特数组、内存地址等。进一步将这些数组织在一起，就可以表达
更多的对象，如数据包、像素点、诗歌，甚至其它任何对象。Go语言提供了丰富的数据组织形式，这依赖于Go语言内置的数据类型。

Go语言将数据类型分为四类：基础类型、复合类型、引用类型和接口类型。基础数据类型包括，数字、字符串、布尔型。

### 整型
Go语言的数值类型包括了几种不同大小的整数、浮点数和复数。每种数值类型都决定了数值的大小范围和是否支持正负符号。Go语言同时提供了有符号和无符号类型
的整数运算。int8、int16、int32、int64分别对应8、16、32、64bit大小的有符号整数，与此对应的uint8、uint16、uint32、uint64对应无符号整数。

Unicode字符rune类型是和int32等价的类型，通常用于表示一个Unicode码点。同样byte也是和uint8类型的等价类型，byte类型一般用于强调数值是一个
原始的数据而不是一个小的整数。

还有一种无符号的整数类型uintptr，没有指定具体的bit大小但是足以容纳指针。uintptr类型只有在底层编程时才需要，特别是Go语言和C语言函数库或操作
系统接口相交互的地方。

### 浮点数
Go语言提供了两种精度的浮点数，float32和float64。浮点数的范围极限值可以在math包找到。常量math.MaxFloat32表示float32能表示的最大数值，
大约是3.4e38；对应的math.MaxFloat64常量大约是1.8e308。它们分别能表示的最小值近似为1.4e-45和4.9e-324。

一个float32类型的浮点数可以提供大约6个十进制数的精度，而float64则可以提供约15个十进制数的精度；通常应该优先使用float64类型，因为float32
类型的累计计算误差很容易扩散，并且float32能精确表示的正整数并不是很大
```
var f float32 = 16777216        // 1 << 24
fmt.Println(f == f+1)           // true
```
浮点数的字面值可以直接写小数部分，像这样：
```
const e = 2.71828
```
使用Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，但是对应表格的数据，使用%e（带指数）或%f的形式打印可能更适合。
```
for x := 0; x < 8; x++ {
    fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
} 
```
函数math.IsNaN用于测试一个数是否为非数NaN，math.NaN则返回非数对应的值。虽然可以用math.NaN来表示一个非法的结果，但是测试一个结果是否是非数
NaN则是充满风险的，因为NaN和任何数都是不相等的。
```
nan := math.NaN()
fmt.Println(nan == nan, nan < nan, nan > nan)       // false, flase, false
```

### 复数
Go语言提供了两种精度的复数类型：complex64和complex128，分别对应float32和float64两种浮点数精度。内置的complex函数用于构建复数，内建的
real和imag函数分别返回复数的实部和虚部：
```
var x complex128 = complex(1, 2)    // 1+2i
var y complex128 = complex(3, 4)    // 3+4i
fmt.Println(x*y)                    // -5+10i
fmt.Println(real(x*y))              // -5
fmt.Println(imag(x*y))              // 10
```
math/cmplx包提供了复数处理的许多函数，例如求复数的平方根函数和求幂函数：
```
fmt.Println(cmplx.Sqrt(-1))     // 0+1i
```
下面的程序使用complex128复数算法生成一个Mandelbrot图像：
[gopl-zh-learn/ch3/mandelbrot](ch3/mandelbrot.go)
```
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
```
![img](ch3/a.png)

### 布尔型
一个布尔类型的值只有两种：true和false。if和for语句的条件部分都是布尔类型的值，并且==和<等比较运算符也会产生布尔型的值。布尔值可以和&&（AND）
和||（OR）操作符结合，并且有短路行为：如果运算符左边值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值，因此下面的表达式总是安全的：
```
s != "" && s[0] == "x"
```
其中s[0]操作如果应用于空字符串将会导致panic异常。

布尔值并不会隐式转换为数字值0和1，反之亦然。必须使用一个显式的if语句辅助转换：
```
i := 0
if b {
    i = 1
}
```
如果经常需要做类似的转换，包装成一个函数会更方便：
```
func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
}
```

### 字符串
一个字符串是一个不可改变的字节序列。字符串可以包含任意的数据，包括byte值0，但是通常是用来包含人类可读的文本。文本字符串通常被解释为采用UTF8编码
的Unicode码点（rune）序列。

内置的len函数可以返回一个字符串中的字节数目（不是rune字符数目），索引操作s[i]返回第i个字节的字节值：
```
s := "hello, world"
fmt.Println(len(s))         // 12
fmt.Println(s[0], s[7])     // 104 119
```
如果试图访问超出字符串索引范围的字节将会导致panic异常：
```
c := s[len(s)]      // panic: index out of range
```
第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个字节。子字符串操作s[i:j]基于原始的s字符串的第ige字节开始到
第j个字节（不包含第j个字节本身）生成一个新的字符串。生成的新字符串将包含j-i个字节:
```
fmt.Println(s[0:5])     // hello
fmt.Println(s[:5])      // world
fmt.Println(s[7:])      // world
fmt.Println(s[:]        // hello, world
```
#### 字符串和Byte切片
标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并
等功能。
* bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。因为字符串是只读的，因此逐步构建字符串会导致大量的内存分配和复制。
在这种情况下，使用bytes.Buffer类型将会更有效。
* strconv包提供了布尔型、整数型、浮点型和对应字符串的相互转换，还提供了双引号转义相关的转换。
* unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。每个函数有一个单一的rune类型的参数，然后返回一个
布尔值。而像ToUpper和ToLower之类的转换函数将用于rune字符的大小写转换。所有的这些函数都是遵循Unicode标准定义的字母、数字等分类规范。strings
包也有类似的函数，它们是ToUpper和ToLower，将原始字符串的每个字符都做相应的转换，然后返回新的字符串。

下面的basename函数实现文件名前缀：
```go
package main

import "fmt"

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {
	fmt.Println(basename("a/b/c.go"))   // c
	fmt.Println(basename("c.d.go"))     // c.d
	fmt.Println(basename("abc"))        // abc
}
```
使用strings库函数实现：
```
func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
```
一个字符串是包含的只读字节数组，一旦创建，是不可变的。相比之下，一个字节slice的元素则可以自由地修改。
字符串和字节slice之间可以相互转换：
```
s := "abc"
b := []byte(s)
s2 := string(b)
```
从概念上讲，一个[]byte(s)转换是分配了一个新的字节数组用于保存字符串数据的拷贝，然后引用这个底层的字节数组。编译器的优化可以避免在一些场景下分配
和复制字符串数据，但总的来说需要确保在变量b被修改的情况下，原始的s字符串也不会改变。将一个字节slice转到字符串的string(b)操作则是构造一个字符串
拷贝，以确保s2字符串是只读的。

为了避免转换中不必要的内存分配，bytes包和strings同时提供了许多实用函数。strings包中的六个函数：
```
func Contains(s, substr string) bool
func Count(s, sep string) int
func Fields(s string) []string
func HasPrefix(s, prefix string) bool
func Index(s, sep string) int
func Join(a []string, sep string) string
```
bytes包中也对应的六个函数：
```
func Contains(b, subslice []byte) bool
func Count(s, sep []byte) int
func Fields(s []byte) [][]byte
func HasPrefix(s prefix []byte) bool
func Index(s, sep []byte) int
func Join(s [][]byte, sep []byte) []byte
```
bytes包还提供了Buffer类型用于字节slice的缓存。一个Buffer开始是空的，但是随着string、byte或[]byte等类型数据的写入可以动态增长，一个
bytes.Buffer变量并不需要初始化，因为零值也是有效的：
```
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intToString([]int{1, 2, 3}))
}

func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
```
#### 字符串和数字的转换
除了字符串、字符、字节之间的转换，字符串和数值之间的转换也比较常见。由strconv包提供这类转换功能。

将一个整数转换为字符串，一种方法是用fmt.Sprintf返回一个格式化的字符串；另一种方法是strconv.Itoa("整数到ASCII")：
```
x := 123
y := fmt.Sprintf("%d", x)
fmt.Println(y, strconv.Itoa(x)) // 123 123
```
FormatInt和FormatUint函数可以用不同的进制来格式化数字：
```
fmt.Println(strconv.FormatInt(int64(x), 2)) // 1111011
```
fmt.Printf函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多，特别是需要在包含附加额外信息的时候：
```
s := fmt.Sprintf("x=%b", x)     // x=1111011
```
如果要将一个字符串解析为整数，可以使用strconv包的Atoi或ParseInt函数，还有用于解析无符号整数的ParseUInt函数：
```
x, err := strconv.Atoi("123")                   // x is an int
y, err := strconv.ParseInt("123", 10, 64)       // base 10, up to 64 bits
```
