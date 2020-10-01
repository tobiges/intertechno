package intertechno

import (
	"errors"
	"sync"

	"github.com/stianeikeland/go-rpio/v4"
)

// ErrIntertechnoManagerClosed when IntertechnoManager is closed
var ErrIntertechnoManagerClosed = errors.New("IntertechnoManager is closed")

// ErrIntertechnoCommandBufferFull command buffer full
var ErrIntertechnoCommandBufferFull = errors.New("Intertechno command buffer full")

const commandAsyncMaxBufferSize = 50

// Manager manages all the rc 433mhz switching operations
// it can be accessed by multiple goroutines concurrently
type Manager struct {
	sync.RWMutex
	pin       rpio.Pin
	closed    bool
	asyncMode bool
	asyncCmds chan Command
}

// NewIntertechnoManager returns a new IntertechnoManager
func NewIntertechnoManager(pin int, async bool) (*Manager, error) {
	if err := rpio.Open(); err != nil {
		return nil, err
	}
	im := &Manager{
		pin:       rpio.Pin(pin),
		asyncMode: async,
	}
	im.pin.Output()
	if async {
		im.asyncCmds = make(chan Command, commandAsyncMaxBufferSize)
		go im.handleAsync()
	}
	return im, nil
}

// Close closes the IntertechnoManager and cleans up
// in async mode it doesn't wait for all commands to send
func (im *Manager) Close() error {
	im.Lock()
	defer im.Unlock()
	im.closed = true
	return rpio.Close()
}

// ExecuteCommand executes the passed command
func (im *Manager) ExecuteCommand(c Command) error {
	if im.asyncMode {
		im.RLock()
		defer im.RUnlock()
	} else {
		im.Lock()
		defer im.Unlock()
	}
	if im.closed {
		return ErrIntertechnoManagerClosed
	}
	if err := c.isValid(); err != nil {
		return err
	}
	if im.asyncMode {
		select {
		case im.asyncCmds <- c:
		default:
			return ErrIntertechnoCommandBufferFull
		}
	} else {
		im.transmit(c)
	}
	return nil
}

func (im *Manager) handleAsync() {
	for c := range im.asyncCmds {
		im.RLock()
		if im.closed {
			return
		}
		im.transmit(c)
		im.RUnlock()
	}
}
