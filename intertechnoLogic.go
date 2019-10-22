package intertechno433mhz

import (
	"time"
)

const (
	repeats         = 15  // [0..8] The 2log-Number of times the signal is repeated. The actual number of repeats will be 2^repeats. 2 would be bare minimum, 4 seems robust, 8 is maximum (and overkill).
	periodusec uint = 260 // Duration of one period, in microseconds. One bit takes 8 periods (but only 4 for 'dim' signal).
)

func (im *IntertechnoManager) transmit(c Command) {
	im.Lock()
	defer im.Unlock()
	for i := repeats; i >= repeats; i-- {
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

func (im *IntertechnoManager) sendStartPulse() {
	im.setPinHigh()
	sleepPeriodusec()
	im.setPinLow()
	time.Sleep(time.Microsecond * time.Duration(periodusec*10+(periodusec>>1))) // Actually 10.5T insteat of 10.44T. Close enough.
}

func (im *IntertechnoManager) sendAddress(address uint) {
	for i := 25; i >= 0; i-- {
		im.sendBit((address>>i)&1 != 0)
	}
}

func (im *IntertechnoManager) sendUnit(unit uint) {
	for i := 3; i >= 0; i-- {
		im.sendBit(unit&(1<<i) != 0)
	}
}

func (im *IntertechnoManager) sendDim(dimvalue uint, unit uint) {
	im.setPinHigh()
	sleepPeriodusec()
	im.setPinLow()
	sleepPeriodusec()
	im.setPinHigh()
	sleepPeriodusec()
	im.setPinLow()
	sleepPeriodusec()

	im.sendUnit(unit)

	for i := 3; i >= 0; i-- {
		im.sendBit(dimvalue&(1<<i) != 0)
	}
}

func (im *IntertechnoManager) sendStopPulse() {
	im.setPinHigh()
	sleepPeriodusec()
	im.setPinLow()
	sleepCustomPeriodusec(40)
}

func (im *IntertechnoManager) sendBit(isBitOne bool) {
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

func sleepCustomPeriodusec(m uint) {
	time.Sleep(time.Microsecond * time.Duration(periodusec*m))
}

func sleepPeriodusec() {
	time.Sleep(time.Microsecond * time.Duration(periodusec))
}

func (im *IntertechnoManager) setPinHigh() {
	im.pin.High()
}

func (im *IntertechnoManager) setPinLow() {
	im.pin.Low()
}
