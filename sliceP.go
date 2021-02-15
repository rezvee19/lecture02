package main
import "fmt"

func main(){
x:= [6]string{"aa","bb","cc","dd","ee","ff"}
s:=x[1:5]
s[1]="changed"
fmt.Println(s[1]);
fmt.Println(x);
}