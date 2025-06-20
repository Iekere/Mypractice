package main
import "fmt"
func Myfunction(x int, y int) (result int) {
result = x * y
return result
}

func main() {
fmt.Println(Myfunction(1, 2))
}
