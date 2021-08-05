package intertechno

import (
	"time"
)

const (
	// [0..8] The 2log-Number of times the signal is repeated.
	// The actual number of repeats will be 2^repeats.
	// 2 would be bare minimum, 4 seems robust, 8 is maximum (and overkill).
	repeats = 4

	// Duration of one period, in microseconds.
	// One bit takes 8 periods (but only 4 for 'dim' signal).
	periodusec = 260

	addressBits  = 26
	unitBits     = 4
	dimvalueBits = 4
)

func (im *Manager) transmit(c Command) {
	for i := 1 << repeats; i > 0; i-- {
		im.sendStartPulse()
		im.sendAddress(c.Address)
		im.sendBit(c.Group)

		if c.Action == ActionDim {
			im.sendDim(c.Dimvalue, c.Unit)
		} else {
			im.sendBit(c.Action != 0)
			im.sendUnit(c.Unit)
		}

		im.sendStopPulse()
	}
}

func (im *Manager) sendStartPulse() {
	im.setPinHigh()
	sleepPeriodusec()
	im.setPinLow()

	// Actually 10.5T insteat of 10.44T. Close enough.
	time.Sleep(time.Microsecond * time.Duration(periodusec*10+(periodusec>>1))) 
}

func (im *Manager) sendAddress(address int) {
	for i := addressBits - 1; i >= 0; i-- {
		im.sendBit((address>>i)&1 != 0)
	}
}

func (im *Manager) sendUnit(unit int) {
	for i := unitBits - 1; i >= 0; i-- {
		im.sendBit(unit&(1<<i) != 0)
	}
}

func (im *Manager) sendDim(dimvalue int, unit int) {
	im.setPinHigh()
	sleepPeriodusec()
	im.setPinLow()
	sleepPeriodusec()
	im.setPinHigh()
	sleepPeriodusec()
	im.setPinLow()
	sleepPeriodusec()

	im.sendUnit(unit)

	for i := dimvalueBits - 1; i >= 0; i-- {
		im.sendBit(dimvalue&(1<<i) != 0)
	}
}

func (im *Manager) sendStopPulse() {
	im.setPinHigh()
	sleepPeriodusec()
	im.setPinLow()
	sleepCustomPeriodusec(40)
}

func (im *Manager) sendBit(isBitOne bool) {
	if isBitOne {
		// Send '1'
		im.setPinHigh()
		sleepPeriodusec()
		im.setPinLow()
		sleepCustomPeriodusec(5)
		im.setPinHigh()
		sleepPeriodusec()
		im.setPinLow()
		sleepPeriodusec()
	} else {
		// Send '0'
		im.setPinHigh()
		sleepPeriodusec()
		im.setPinLow()
		sleepPeriodusec()
		im.setPinHigh()
		sleepPeriodusec()
		im.setPinLow()
		sleepCustomPeriodusec(5)
	}
}

func sleepCustomPeriodusec(m int) {
	time.Sleep(time.Microsecond * time.Duration(periodusec*m))
}

func sleepPeriodusec() {
	time.Sleep(time.Microsecond * time.Duration(periodusec))
}

func (im *Manager) setPinHigh() {
	im.pin.High()
}

func (im *Manager) setPinLow() {
	im.pin.Low()
}
