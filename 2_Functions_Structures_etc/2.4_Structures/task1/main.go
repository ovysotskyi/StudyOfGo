package main

import "fmt"

type Bot struct {
	On    bool
	Ammo  int
	Power int
}

func (bot *Bot) Shoot() bool {
	if bot.On && bot.Ammo > 0 {
		bot.Ammo--
		return true
	}

	return false
}

func (bot *Bot) RideBike() bool {
	if bot.On && bot.Power > 0 {
		bot.Power--
		return true
	}

	return false
}

func main() {
	testStruct := &Bot{}

	//debug

	testStruct.On = true
	testStruct.Ammo = 3
	testStruct.Power = 4

	fmt.Print(testStruct)

	testStruct.Shoot()

	fmt.Print(testStruct)

}
