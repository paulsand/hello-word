/*package main
import (
	"bytes"
	"errors"
	"fmt"
	"math"
)

var errDivisonByZero = errors.New("Divided by 0")
// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() (int,error) {
	// 返回一个闭包
	return func() (int, error) {
		if(value==0)
		return 0,errDivisonByZero
		// 累加
		value++
		// 返回一个累加值
		return value, nil
	}
}
func main() {
	// 创建一个累加器, 初始值为1
	accumulator := Accumulate(1)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)
	// 创建一个累加器, 初始值为1
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator2)
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)
	accumulator3 := Accumulate(130)
	// 累加1并打印
	fmt.Println(accumulator3())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator3)

	fmt.Printf("pure string\n")
	fmt.Printf("value: %v %f\n", true, math.Pi)
	fmt.Println(joinStrings("hammer", " mom", " and", 10,true))

}


func joinStrings(slist ...interface{}) string {
	// 定义一个字节缓冲, 快速地连接字符串
	//var b bytes.Buffer
	var ty bytes.Buffer
	//var c bytes.Buffer
	// 遍历可变参数列表slist, 类型为[]string
	for _, s := range slist {
		// 将遍历出的字符串连续写入字节数组
		switch  s.(type){
		case bool :
			ty.WriteString("bool   ")
		case string:   // 当s为字符串类型时
			ty.WriteString("string   ")
		case int:    // 当s为整型类型时
			ty.WriteString("int   ")
		}
	}
	// 将连接好的字节数组转换为字符串并输出
	return ty.String()

}*/
/*
package  main

import (
	"fmt"
)

type ParseError struct{
	Filename string
	Line int
}

func(e *ParseError) Error() string{
	return fmt.Sprintf("s%:d%",e.Filename,e.Line)


}
func newParseError(filename string, line int) error{
	return  &ParseError{filename,line}
}

func main(){
	var e error
	e=newParseError("main.go", 1)
	fmt.Println(e.Error())
	switch detail :=e.(type) {
	case *ParseError:
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default: // 其他类型的错误
		fmt.Println("other error")
	}


}
*/
/*
package main
import (
	"fmt"
	"runtime"
)
// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}
// 保护方式允许一个函数
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()
	entry()
}
func main() {
	fmt.Println("运行前")
	// 允许一段手动触发的错误
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		// 使用panic传递上下文
		panic(&panicContext{
			"手动触发panic",
		})
		fmt.Println("手动宕机后")
	})
	// 故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
}
//testst
*/

/*

package main
import "fmt"
// 车轮
type Wheel struct {
	Size int
}
// 车
type Car struct {
	Wheel Wheel
	// 引擎
	Engine struct {
		Power int    // 功率
		Type  string // 类型
	}
}
func main() {
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化引擎
		Engine: struct {
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
}

*/
/*

package main
import (
	"fmt"
	"runtime"
	"time"
)
func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}
func main() {


//	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	// 并发执行程序
	go running()
	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Println("input")
	fmt.Scanln(&input)
	fmt.Println(input)
	fmt.Scanln(&input)
}

*/

/*
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
*/
/*
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func server(address string, exitChan chan int) {

	l,err := net.Listen("tcp",address)

	if err!=nil{
		fmt.Println(err.Error())
		exitChan<-1
	}

	fmt.Println("listen: "+ address)
	defer l.Close()

	for{
		con,err :=l.Accept()
		if err!=nil{
			fmt.Println(err.Error())
			continue
		}
		go handleSession(con,exitChan)

	}
}

func handleSession(con net.Conn, exitChan chan int){

	reader := bufio.NewReader(con)
	for{
		str,err:=reader.ReadString('\n')
		if err==nil{
			str = strings.TrimSpace(str)
			if !processTelnetCommand(str,exitChan){
				con.Close()
				break
			}
			con.Write([]byte(str+"\r\n"))
		}else{
			fmt.Println("Session closed")
			con.Close()
			break
		}


	}



} */
/*
func processTelnetCommand(str string, exitChan chan int ) bool{
	if strings.HasPrefix(str,"@close"){
		fmt.Println("Session closed")
		// 告诉外部需要断开连接
		return false
	}else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		exitChan<-0
		return false
	}
	fmt.Println(str)
	return true
}

func main(){

	address := "127.0.0.1:7001"
	exitChan:= make(chan int)
	go server(address , exitChan)

	code := <-exitChan
	os.Exit(code)


} */

/*
package main
import (
	"fmt"
	"reflect"
)
func main() {
	// 声明一个空结构体
	type cat struct {
		Name string `name:"typename" id:"123"ok:"afda"`
		// 带有结构体tag的字段
		Type int `json:"type" id:"100"`
		Type1 int `jsons:"type1" id:"110"`
	}
	// 创建cat的实例
	ins := cat{Name: "mimi", Type: 1}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Name"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("ok"), catType.Tag.Get("id"),"name")
		fmt.Println(catType.Tag.Get("name"), catType.Tag.Get("id"))
	}


	// 声明整型变量a并赋初值
	var a int = 1024
	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)
	// 获取interface{}类型的值, 通过类型断言转换
	value12:= reflect.ValueOf(a)

	getInterface := value12.Interface().(int)





	var getA int = valueOfA.Interface().(int)
	// 获取64位的值, 强制类型转换为int类型
	var getA2 int64 = (valueOfA.Int())
	fmt.Println(getA, getA2, getInterface)
}

*/

/*
package main
import (
	"fmt"
	"reflect"
)
// 定义结构体
type dummy struct {
	a int
	b string
	// 嵌入字段
	float32
	bool
	next *dummy
}
func main() {
	// 值包装结构体
	d := reflect.ValueOf(dummy{
		next: &dummy{ b:"string2",},
	})
	// 获取字段数量
	fmt.Println("NumField", d.NumField(),)
	// 获取索引为2的字段(float32字段)
	floatField := d.Field(2)
	// 输出字段类型
	fmt.Println("Field", floatField.Type())
	// 根据名字查找字段
	fmt.Println("FieldByName(\"b\").Type", d.FieldByName("b").Type())
	// 根据索引查找值中, next字段的int字段的值
	fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 0}).Type(),"FieldByIndex([]int{4, 1}) value",d.FieldByIndex([]int{4, 1}) )
}

*/


package main
import (
	"fmt"
	"reflect"
)
func main() {
	type dog struct {
		LegCount int
	}
	// 获取dog实例地址的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})
	// 取出dog实例地址的元素
	valueOfDog = valueOfDog.Elem()
	dogcolon := reflect.New(valueOfDog.Type())
	println("dogcolong", valueOfDog.Type(),dogcolon.Type())

//	fmt.Println(valueOfDog.Int())
	// 获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("LegCount")

	// 尝试设置legCount的值(这里会发生崩溃)
	vLegCount.SetInt(4)
	fmt.Println(vLegCount.Int())


	var A int
	// 取变量a的反射类型对象
	typeOfA := reflect.TypeOf(A)
	fmt.Println("typeofA:", typeOfA)
	// 根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfA)
	valueOfa := aIns.Elem()
	valueOfa.Set(reflect.ValueOf(10))
	// 输出Value的类型和种类
	fmt.Println(aIns.Type(), aIns.Kind(),valueOfa.Int())

	funcValue :=reflect.ValueOf(add)

	parameter := [] reflect.Value{reflect.ValueOf(10),reflect.ValueOf(20)}

	retValue :=funcValue.Call(parameter)
	println(retValue[0].Int())

}


func add(a, b int) int {
	return a + b
}
