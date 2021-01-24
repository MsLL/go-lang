package main

import (
	. "../lang-base"
)

func main() {
	var beheave Beheave
	var beheaveCat = new(Cat)
	var beheaveDog = new(Dog)

	beheave = beheaveCat
	beheave.Eat()
	beheave.Sleep()

	beheave = beheaveDog
	beheave.Eat()
	beheave.Sleep()
}
