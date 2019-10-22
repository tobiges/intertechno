package main

import (
	"github.com/tobiges/intertechno433mhz"
)

const pin = 6

func main() {
	intertechnoManager, err := intertechno433mhz.NewIntertechnoManager(pin)
	if err != nil {
		panic(err)
	}
	defer intertechnoManager.Close()
	// Turn Device with address 45234543 off
	c := intertechno433mhz.Command{
		Address: 45234543,
		Action:  intertechno433mhz.ActionOff,
	}
	intertechnoManager.ExecuteCommand(c)

	// Dim Device with address 46256432 and unit 4 to the dimvalue 7
	c = intertechno433mhz.Command{
		Address:  46256432,
		Action:   intertechno433mhz.ActionDim,
		Dimvalue: 7,
		Unit:     4,
	}
	intertechnoManager.ExecuteCommand(c)

	// Turn Group with address 425452345 on
	c = intertechno433mhz.Command{
		Address: 425452345,
		Action:  intertechno433mhz.ActionOn,
		Group:   true,
	}
	intertechnoManager.ExecuteCommand(c)

	intertechnoManager.Close()
}
