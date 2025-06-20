package main
import "fmt"
func Myfunction(x int, y string) (result int, txt string) {
result = x + x
txt = y + " Estate"
return
}
func main() {
fmt.Println(Myfunction(12, "Eskay"))
}
