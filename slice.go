package main
import "fmt"
//import "reflect"

func main(){
//students := [...] string {"Rubel", "Hassan", "Anik", "Karim", "Ovi", "Jahin"}

//sliced:= students[0:4]
//var fruits []string
//fruits= append(fruits,"Apple", "Banana", "Orange","Sapoda")

//fmt.Println(fruits)
//fmt.Printf("%T \n",fruits)
//fmt.Printf("%T \n",students)
//x := reflect.TypeOf(fruits).Kind().String()
//fmt.Println(x)
var mapName map[string]int

mapName = map[string]int{
    "rsc": 3711,
    "r":   2138,
    "gri": 1908,
    "adg": 912,
}
fmt.Println(mapName)

}
