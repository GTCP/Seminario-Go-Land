package main 

import "fmt"

//Printable ..
type Printable interface{
	print()
}
type person struct{
	name string
}
type figure struct{
	name string
}
func (f *figure) print(){
	fmt.Println("[figure]",f.name)
} 

func (p *person) print(){
	fmt.Println("[person]",p.name)
}
func invokePrint(p Printable){
	p.print()
}
//func (p *person) clean(){
//p.name = ""
//}
func main(){
	p := &person{name: "juan"}
	invokePrint(p)
	f := &figure{name: "Cube"}
 }