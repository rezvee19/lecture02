package main
import "fmt"

var x string= "Hi, I am variable, any dunction can use me"

func main(){
fmt.Println(x)
f()
}

func f(){
fmt.Println(x)
}