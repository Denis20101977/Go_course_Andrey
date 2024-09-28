package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight float64
	height float64
}

func (u User) PrintUserInfo() {
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}

func NewUser(name, sex string, age int, weight, height float64) User {
	return User{
		name:   name,
		age:    age,
		sex:    sex,
		weight: weight,
		height: height,
	}
}

func main() {
	user := NewUser("Vasya", "Male", 25, 188.5, 196.5)
	user1 := NewUser("Vasya", "Male", 125, 1188.5, 1196.5)

	user.PrintUserInfo()
	user1.PrintUserInfo()

}
