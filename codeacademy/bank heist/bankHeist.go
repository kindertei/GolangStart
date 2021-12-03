package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(40)
	if eludedGuards <= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job? but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	fmt.Println(isHeistOn)
	openedVault := rand.Intn(100)
	if isHeistOn && openedVault > 40 {
		fmt.Println("Grab and GO!")
	} else if !isHeistOn || openedVault <= 40 {
		isHeistOn = false
		fmt.Println("Fuck you")
	}
	leftSafety := rand.Intn(100)
	if isHeistOn {
		switch leftSafety {
		case 0:
			isHeistOn = false
			fmt.Println("Fail leftSafety")
		case 1:
			isHeistOn = false
			fmt.Println(" Fucking door")
		default:
			fmt.Println("Success exit")
		}
		if isHeistOn {
			amtStolen := 1000 + rand.Intn(100000)
			fmt.Println(amtStolen)
		}
	}

}
