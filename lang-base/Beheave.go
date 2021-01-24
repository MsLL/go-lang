package lang_base

import "fmt"

type Dog struct {
}

func (dog *Dog) Eat() {
	fmt.Println("dog eat")
}
func (dog *Dog) Sleep() {
	fmt.Println("dog sleep")

}

type Cat struct {
}

func (cat *Cat) Eat() {
	fmt.Println("cat eat")

}
func (cat *Cat) Sleep() {
	fmt.Println("cat sleep")
}
