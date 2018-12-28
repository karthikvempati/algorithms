package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"github.com/karthikvempati/algorithms/gotour"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
var c, python, java bool

func main() {
	//a, b := swap("karthik", "hello")
	//fmt.Println(a, b)
	//fmt.Println(split(34))
	//random()
	//basictypes()
	//zerovalues()
	/*
	defer WithDefer()
	fmt.Println("After defer")
	NoDefer()
	a := []int{1,2,3,4,8}
	Sum(append(a[1:3],5))

	value, even := HalfOddEven(9)
	fmt.Println(value, even)

	value1, even1 := HalfOddEven(123132)
	fmt.Println(value1, even1)

	fmt.Println(MaxValue(1,2,3,4,5,6,7,8,9))
	EvenGen := makeEvenGenerator()
	fmt.Println(EvenGen())
	fmt.Println(EvenGen())
	fmt.Println(EvenGen())
	fmt.Println(EvenGen())

	OddGen := makeOddGenerator()
	fmt.Println(OddGen())
	fmt.Println(OddGen())
	fmt.Println(OddGen())
	fmt.Println(OddGen())


	x,y := 1,2
	fmt.Println(swap1(&x,&y))
	*/
	c := gotour.Circle{Radius:9}
	r := gotour.Rectangle{Width:8,Breadth:6}
	printArea(c, r)
	printCircumference(c,r)
}

func printArea(s...gotour.Shape){
	for _,a := range s{
		fmt.Println(a.Area())
	}
}

func printCircumference(s...gotour.Shape){
	for _,a := range s{
		fmt.Println(a.Circumference())
	}
}

func makeEvenGenerator() func() uint64 {
	i := uint64(0)
	return func() uint64 {
		i = i + 2
		return i
	}
}

func makeOddGenerator() func() uint64{
	i := uint64(1)
	firstCall := true
	return func() uint64 {
		if firstCall {
			firstCall = false
			return i
		}
		i = i+ 2
		return i
	}
}

func swap1(x *int,y *int) (int, int){
	return *y,*x
}

func Sum(array []int){
	s := 0
	for _, value := range array {
		s += value
	}

	fmt.Println(s);
}

func MaxValue(a...int) int{
	max := 0

	for _,val := range a {
		if max < val{
			max = val
		}
	}

	return max
}

func HalfOddEven(i int) (int,bool) {
	defer func ()  {
		str := recover()
		if(str != nil){
			fmt.Println(str)
		}
	}()
	if(i > 10){
		panic("integer overflows")
	}
	return (i+2147483647)/2, i%2==0
}

func random() {
	var i int
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(i, c, python, java)
}

func swap(x, y string) (string, string) {
	return y, x
}

func NoDefer(){
	fmt.Println("No differ")
}

func WithDefer()  {
	fmt.Println("With Defer")
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func basictypes() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func zerovalues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
