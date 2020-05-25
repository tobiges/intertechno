package main

import (
	"github.com/tobiges/intertechno"
)

const pin = 6

func main() {
	intertechnoManager, err := intertechno.NewIntertechnoManager(pin)
	if err != nil {
		panic(err)
	}
	defer intertechnoManager.Close()
	// Turn Device with address 45234543 off
	c := intertechno.Command{
		Address: 45234543,
		Action:  intertechno.ActionOff,
	}
	intertechnoManager.ExecuteCommand(c)

	// Dim Device with address 46256432 and unit 4 to the dimvalue 7
	c = intertechno.Command{
		Address:  46256432,
		Action:   intertechno.ActionDim,
		Dimvalue: 7,
		Unit:     4,
	}
	intertechnoManager.ExecuteCommand(c)

	// Turn Group with address 425452345 on
	c = intertechno.Command{
		Address: 425452345,
		Action:  intertechno.ActionOn,
		Group:   true,
	}
	intertechnoManager.ExecuteCommand(c)

	intertechnoManager.Close()
}
