package main 

import "fmt" 

func main(){
    var arr[2]int 
    arr[0] = 1
	arr[1] = 2

    for i,v := range arr {
        fmt.Println(i, v)
    }
	fmt.Println("--------")

	l := make( []int,3)
	l = append(l, 10)
	fmt.Println("%p\n, l")
	l = append(l, 100)
	fmt.Println("%p\n, l")
	l = append(l, 1000)	
	fmt.Println("%p\n, l")
	for i,v := range l {
		fmt.Println(i, v)
	}
	fmt.Println("---------------")

	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"
	for k,v := range m{
		fmt.Println(k, v)
	}
}