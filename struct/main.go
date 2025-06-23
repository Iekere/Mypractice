package main
import "fmt"

type Person struct {
name string
age int
job string
salary int
}
func main() {
var pers1 Person
var pers2 Person

pers1.name = "Innocent"
pers1.age = 27
pers1.job = "Software Engineer"
pers1.salary = 8500000
pers2.name = "Esther"
pers2.age = 10
pers2.job = "Doctor"
pers2.salary = 3500000

fmt.Println("Name: ", pers1.name)
fmt.Println("Age: ", pers1.age)
fmt.Println("Job: ", pers1.job)
fmt.Println("Salary: ", pers1.salary)
fmt.Println("Name: ", pers2.name)
fmt.Println("Age: ", pers2.age)
fmt.Println("Job: ", pers2.job)
fmt.Println("Salary: ", pers2.salary)
}
