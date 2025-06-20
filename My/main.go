package main
import "fmt"
func Myfunction(x int, y string) (result int, txt string) {
result = x + x
txt = y +" Estate"
return
}
func main() {
a, _ := Myfunction(12, "Eskay")
fmt.Println(a)
}
