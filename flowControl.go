package main
import ( 
    "fmt"
)
func main(){
fmt.Println("Enter your age:")
var age int
fmt.Scanf("%d",&age)
fmt.Println("your age is","hello",age)
var s string = "hi"
var x string = "hello"
fmt.Println(s,x) //comma adds a "space"
fmt.Println(s+x) //+ concateate doesnt add space
}
