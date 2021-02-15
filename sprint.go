package main
import "fmt"

func main(){
	x:=fmt.Sprintln("hello world")
	fmt.Println(x)
	b:=foo()
	fmt.Println(b)
}

func foo() string{
var name string
fmt.Scanf("%s",&name)
return fmt.Sprintf("here is his name, %s",name)
}