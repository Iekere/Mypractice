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
pers1.salary = 85000000
pers2.name = "Gloria"
pers2.age = 10
pers2.job = "Doctor"
pers2.salary = 5800000

PrintPerson(pers1)
PrintPerson(pers2)
}
func PrintPerson(pers Person) {
fmt.Println("Name: ", pers.name)
fmt.Println("Age: ", pers.age)
fmt.Println("Job: ", pers.job)
fmt.Println("Salary: ", pers.salary)
}
