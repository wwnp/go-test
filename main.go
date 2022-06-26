package main

import "fmt"

//------------------------------------------
// TYPES
type Age int

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

//------------------------------------------
// METHODS
func (u User) printUserInfo() {
	fmt.Println("printUserInfo: ", u.name, u.age, u.sex) //  по value- копируется структура
} // метод структуры по value; (u User) - ресивер

func (u User) getName() string {
	return "getName: " + u.name
} // метод структуры просто вернуть значение

func (u *User) setName(newName string) {
	fmt.Println("setName on:", newName)
	u.name = newName
} // метод структуры по pointer(ссылке) то исп та же структура

func (a Age) isAdult() string {
	if a >= 18 {
		return "Can pass"
	}
	return "Can't pass "
} // Метод для Type Age

//------------------------------------------
// CONSTRUCTOR
func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    Age(age),
		height: height,
		weight: weight,
	}
} // constructor

//------------------------------------------
func main() {
	user := NewUser("Pole", "male", 19, 67, 185)

	user.printUserInfo()

	fmt.Println(user.getName())

	user.setName("YANHEN")

	fmt.Println(user.getName())

	fmt.Println(user.age.isAdult())
}
