//main包才可以编译为可执行文件
package main

import (
	//"demo7"
	"demo9"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	//"testing"
)

func init() {
	fmt.Println("init()...")
}

func main() {
	//public函数需要首字母大写
	//go语言不支持方法重载
	//同一个包下不允许出现相同的函数名
	//num := demo2.Calculator(1.2,1.3,'+')
	//当函数有多个返回值，可以用_作为占位符忽略返回值
	/*
		num2,_ := demo2.MultyReturn()
		fmt.Println(num)
		fmt.Println(num2)
	*/
	//函数形参变量名后加...可以实现可变参数，此时形参类似于数组
	//可变参数需要放在形参列表最后
	//函数本身也是一种数据类型可以复制给变量,因此函数也可以作为形参
	/*
		func NtypeFunc(newFunc func(int,int)int,n1 int ,n2 int){
			return newFunc(n1,n2)
		}
	*/
	/*
		newNum := demo2.Num
		fmt.Println(newNum(1,2))
	*/
	//go可以自定义数据类型，自定义数据类型与原数据类型为不同数据类型
	//同时go也支持对返回值命名
	/*
		func NtypeFunc(n1 int, n2 int)(sum int,sub int){
			sum = n1+n2
			sub = n1-n2
			return
		}
	*/

	//map操作注意一定要先make分配内存再使用
	/*
		my_map := make(map[int]string)
		demo3.InitMyMap(my_map)
		demo3.DeleteMap(my_map,0)
		fmt.Println(my_map)
		demo3.FindKeyInMap(my_map,1)
		demo3.FindKeyInMap(my_map,0)
		demo3.ReadMap(my_map)
		demo3.InitMyMap(my_map)
		demo3.SlipOfMap(my_map)
	*/
	/*
		结构体中的字段如果是引用，在赋值时候需要先make
		结构体struct1 := struct2 属于值拷贝，两个属于不同对象(独立内存空间)
		结构体的字段在内存中是连续分布的(利于数据查找)
		结构体实例化过程中可以通过在字段末尾添加tag实现序列化
	*/
	// demo3.InitStruct()
	// fmt.Printf("\n")
	// demo3.InitInstance()
	// a := [...]int{0, 1, 2, 3, 4, 5}
	// demo4.ReverseSlice(a[:2])
	// fmt.Println(a)
	// demo4.ReverseSlice(a[2:])
	// fmt.Println(a)
	// demo4.ReverseSlice(a[:])
	// fmt.Println(a)
	// if strings.HasPrefix("abcde", "abc") {
	// 	fmt.Println("yep")
	// }

	//结构体方法的值传递
	// var cat demo3.Cat
	// cat.Age = 1
	// cat.Color = "white"
	// cat.Name = "babe"
	// cat.Speed = 2.0
	// cat.Eat()
	// fmt.Println(cat.Run(3.2))

	//指针类型的方法
	// node := new(demo3.IntLinkList)
	// head := node
	// for i := 0; i < 3; i++ {
	// 	node.Value = i
	// 	node.Tail = new(demo3.IntLinkList)
	// 	node = node.Tail
	// }
	// fmt.Println(head.Sum())

	//已有类型方法
	// var num4 demo3.Integer = 10
	// fmt.Println(num4.Increace())

	// var stu = demo5.NewStudent("tom", 3, "21927055")
	// fmt.Println(stu)

	//init
	// fmt.Println("main()...")
	// f := demo5.My_closed_func()
	// fmt.Println(f(2))
	// fmt.Println(f(2))

	//closed_func
	// f2 := demo5.MakeSufix(".jpg")
	// fmt.Println(f2("mypicture"))
	// fmt.Println(f2("mypic.jpg"))

	//defer
	//demo5.Func_defer(1, 2, 3, 4, 5)

	//封装
	// person := demo6.NewPerson("zhangsan")
	// person.SetAge(20)
	// person.SetSalary(4000)
	// fmt.Println(person)

	//继承（不建议使用多重继承）
	//如果继承了多个匿名结构体，必须指明是哪个结构体的变量
	//如果嵌套一个有名结构体，则必须指明是哪个结构体变量
	// pupil := new(demo6.Pupil)
	// granduate := new(demo6.Granduate)
	// pupil.Name = "tom"
	// pupil.SetAge(8)
	// pupil.SetScore(98.8)
	// fmt.Println(pupil.ShowInfo())
	// granduate.Name = "holly"
	// granduate.SetAge(24)
	// granduate.SetScore(80.9)
	// fmt.Println(granduate.ShowInfo())

	//接口
	// children := new(demo6.Children)
	// adult := new(demo6.Adult)
	// demo6.Recieve(children)
	// demo6.Recieve(adult)

	//利用接口实现结构体排序
	// var heros demo6.HeroSlice
	// for i := 0; i < 10; i++ {
	// 	hero := demo6.Hero{
	// 		Name:  "tom" + strconv.Itoa(i),
	// 		Score: float64(rand.Intn(100)),
	// 	}
	// 	heros = append(heros, hero)
	// }
	// fmt.Println("排序前...")
	// for _, v := range heros {

	// 	fmt.Println(v)
	// }
	// fmt.Println("排序后...")
	// sort.Sort(&heros)
	// for _, v := range heros {
	// 	fmt.Println(v)
	// }

	//统计文件中字符个数
	//demo6.CountChar("src/demo6/test.txt")
	num := runtime.NumCPU()
	fmt.Println("cpu number: " + strconv.Itoa(num))
	numChan := make(chan int, 100)
	primeChan := make(chan int, 500)
	wg := &(sync.WaitGroup{})
	wg.Add(5)
	go demo9.PutRoutine(500, numChan, wg)
	//time.Sleep(time.Second)
	for i := 0; i < 4; i++ {
		go demo9.IsPrimeRoutine(numChan, primeChan, wg)
		//time.Sleep(time.Second)
	}
	wg.Wait()
	close(primeChan)
	wg.Add(1)
	go demo9.GetPrime(primeChan, wg)
	//time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("执行完成")
}

//单元测试
// func TestAddUpper(t *testing.T) {
// 	res := demo7.AddUpper(10)
// 	if res != 55 {
// 		t.Fatalf("执行错误，结果为%v\n", res)
// 	}
// 	t.Logf("执行正确")
// }
