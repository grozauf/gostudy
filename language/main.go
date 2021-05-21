package main

import "fmt"

type HandlerFunc = func (int, int)int

func sum(a int, b int) int {
	return a + b
}

func Logging(fn HandlerFunc) HandlerFunc {
  return func(a int, b int) int {
    fmt.Println(a, b)
    res := fn(a, b)
    fmt.Println(res)

	return res
  }
}

func main() {
	for i,j := 0,0; i < 10 && j < 10; func (){i += 1; j += 1}() {
		fmt.Println(i, j)
	}

	Logging(sum)(10,20)

	array := []int{1,2,3,4,5,6,7,8}
	var vecArray [][]int

	vecArray = append(vecArray, array[0:2])
	vecArray = append(vecArray, array[2:4])

	fmt.Println(array)
	fmt.Println(vecArray)

	items := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
	  items = append(items, i+1)
	}
	fmt.Println(items)
}
