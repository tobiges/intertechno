package intertechno

import (
	"errors"
	"sync"

	"github.com/stianeikeland/go-rpio/v4"
)

// ErrIntertechnoManagerClosed when IntertechnoManager is closed
var ErrIntertechnoManagerClosed = errors.New("IntertechnoManager is closed")

// Manager manages all the rc 433mhz switching operations
// it can be accessed by multiple goroutines concurrently
type Manager struct {
	sync.Mutex
	pin    rpio.Pin
	closed bool
}

// NewIntertechnoManager returns a new IntertechnoManager
func NewIntertechnoManager(pin int) (*Manager, error) {
	err := rpio.Open()
	if err != nil {
		return nil, err
	}
	im := &Manager{pin: rpio.Pin(pin)}
	im.pin.Output()
	return im, nil
}

// Close closes the IntertechnoManager and cleans up
func (im *Manager) Close() {
	im.Lock()
	defer im.Unlock()
	im.closed = true
	rpio.Close()
}

// ExecuteCommand executes the passed command
func (im *Manager) ExecuteCommand(c Command) error {
	if im.closed {
		return ErrIntertechnoManagerClosed
	}
	if err := c.isValid(); err != nil {
		return err
	}
	if c.Async {
		go im.transmit(c)
	} else {
		im.transmit(c)
	}
	return nil
}
