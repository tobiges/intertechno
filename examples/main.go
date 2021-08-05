package main

import (
	"github.com/tobiges/intertechno"
)

// 433mhz transmitter pin
const pin = 6

func main() {
	intertechnoManager, err := intertechno.NewIntertechnoManager(pin, false)
	if err != nil {
		panic(err)
	}
	defer intertechnoManager.Close()

	// Turn Device with address 45234543 off
	c := intertechno.Command{
		Address: 45234543,
		Action:  intertechno.ActionOff,
	}
	if err := intertechnoManager.ExecuteCommand(c); err != nil {
		panic(err)
	}

	// Dim Device with address 46256432 and unit 4 to the dimvalue 7
	c = intertechno.Command{
		Address:  46256432,
		Action:   intertechno.ActionDim,
		Dimvalue: 7,
		Unit:     4,
	}
	if err := intertechnoManager.ExecuteCommand(c); err != nil {
		panic(err)
	}

	// Turn Group with address 425452345 on
	c = intertechno.Command{
		Address: 425452345,
		Action:  intertechno.ActionOn,
		Group:   true,
	}
	if err := intertechnoManager.ExecuteCommand(c); err != nil {
		panic(err)
	}
}
