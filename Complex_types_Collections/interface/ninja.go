package main

import "fmt"

type ninjaStar struct {
	owner string
}

func (ninjaStar) attack() {
	fmt.Println("Throwing ninja start")
}

type ninjaSword struct {
	owner string
}

func (ninjaSword) attack() {
	fmt.Println("Swinging ninja sword")
}

type ninjaWeapon interface {
	attack()
}

func attack(weapon ninjaWeapon) {
	weapon.attack()
}

func main() {
	stars := []ninjaStar{{owner: "Wallace"}, {owner: "Wallace"}}
	for _, star := range stars {
		star.attack()
	}

	swords := []ninjaSword{{owner: "Wallace"}, {owner: "Wallace"}}
	for _, sword := range swords {
		sword.attack()
	}

	fmt.Println()

	weapons := []ninjaWeapon{ninjaStar{owner: "Wallace"}, ninjaSword{owner: "Wallace"}}
	for _, weapon := range weapons {
		attack(weapon)
	}
}
