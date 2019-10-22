package intertechno433mhz

import (
	"errors"
	"sync"

	"github.com/warthog618/gpio"
)

// ErrIntertechnoManagerClosed when IntertechnoManager is closed
var ErrIntertechnoManagerClosed = errors.New("IntertechnoManager is closed")

// IntertechnoManager manages all the rc 433mhz switching operations
// it can be accessed by multiple goroutines concurrently
type IntertechnoManager struct {
	sync.Mutex
	pin    *gpio.Pin
	closed bool
}

// NewIntertechnoManager returns a new IntertechnoManager
func NewIntertechnoManager(pin int) (*IntertechnoManager, error) {
	err := gpio.Open()
	if err != nil {
		return nil, err
	}
	im := &IntertechnoManager{pin: gpio.NewPin(pin)}
	im.pin.Output()
	return im, nil
}

// Close closes the IntertechnoManager and cleans up
func (im *IntertechnoManager) Close() {
	im.Lock()
	defer im.Unlock()
	im.closed = true
	gpio.Close()
}

// ExecuteCommand executes the passed command
func (im *IntertechnoManager) ExecuteCommand(c Command) error {
	if im.isClosed() {
		return ErrIntertechnoManagerClosed
	}
	if err := c.isValid(); err != nil {
		return err
	}
	im.transmit(c)
	return nil
}

func (im *IntertechnoManager) isClosed() bool {
	im.Lock()
	defer im.Unlock()
	return im.closed
}
