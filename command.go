package intertechno

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidAddress is returned if address is invalid
	ErrInvalidAddress = errors.New("address has to be in range of [0..2^26-1]")
	// ErrInvalidAction is returned if action is invalid
	ErrInvalidAction = errors.New("action has to be ActoinOff (0) => off, ActionOn (1) => on, ActionDim (2) => dim (if device is off it will be turned on and dimmed)")
	// ErrInvalidDimvalue is returned if dimvalue is invalid
	ErrInvalidDimvalue = errors.New("dimvalue has to be in range of [1..15] if action is set to ActionDim (2) -- if device is off it will be turned on and dimmed")
	// ErrInvalidUnit is returned if unit is invalid
	ErrInvalidUnit = errors.New("unit has to be in range of [0..15]")
)

// Action can be ActoinOff (0) => off, ActionOn (1) => on, ActionDim (2) => dim (if device is off it will be turned on and dimmed)
type Action uint

const (
	// ActionOff (0) => off
	ActionOff Action = iota
	// ActionOn (1) => on
	ActionOn
	// ActionDim (2) => dim (if device is off it will be turned on and dimmed)
	ActionDim
	actionEnd
)

func (ia Action) String() string {
	names := [...]string{"ActoinOff", "ActionOn", "ActionDim"}
	if !ia.isValid() {
		return fmt.Sprintf("Invalid Action: %d", ia)
	}
	return names[ia]
}

func (ia Action) isValid() bool {
	return ia < actionEnd
}

// Command is used to store the information to send
type Command struct {
	// Address Address of this transmitter [0..2^26-1]
	Address int
	// Dimvalue [1..15] Dim level if action is set to ActionDim. 15 for brightest level.
	Dimvalue int
	// Action ActoinOff (0) => off, ActionOn (1) => on, ActionDim (2) => dim (dimvalue has to be set)
	Action Action
	// Unit [0..15] unit of the device
	Unit int
	// Group True to send command to the address group.
	Group bool
}

func (c Command) isValid() error {
	if c.Address < 0 || (c.Address>>addressBits) != 0 {
		return ErrInvalidAddress
	} else if !c.Action.isValid() {
		return ErrInvalidAction
	} else if c.Action == ActionDim && (c.Dimvalue < 1 || (c.Dimvalue>>dimvalueBits) != 0) {
		return ErrInvalidDimvalue
	} else if c.Unit < 0 || (c.Unit>>unitBits) != 0 {
		return ErrInvalidUnit
	}
	return nil
}
