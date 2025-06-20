package main
import "fmt"
func Factorial_Recursion(x float64)(y float64) {
if x > 0 {
y = x * Factorial_Recursion(x-1)
} else {
y = 1
}
return
}
func main() {
fmt.Println(Factorial_Recursion(5))
}
