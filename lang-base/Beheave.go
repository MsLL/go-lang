package lang_base

import "fmt"

type Dog struct {
}


//declare struct Dog's method 
func (dog *Dog) Eat() {
	fmt.Println("dog eat")
}
//declare struct Dog's method 
func (dog *Dog) Sleep() {
	fmt.Println("dog sleep")

}

type Cat struct {
}
//declare struct Cat's method
func (cat *Cat) Eat() {
	fmt.Println("cat eat")

}
//declare struct Cat's method
func (cat *Cat) Sleep() {
	fmt.Println("cat sleep")
}
