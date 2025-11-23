/**************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@license	Copyright © 2021-2026 Michael Roberts

/**************************************************************************************/

package adafruit

/**************************************************************************************/

import (
	"machine"
)

/**************************************************************************************/

const DefaultTMP117I2CAddress uint16 = 0x48

/**************************************************************************************/

// NewTMP117Device creates a new Device instance for TMP117 sensor using the default I2C
// address and the specified bus.
func NewTMP117Device(bus machine.I2C) *Device {
	return &Device{
		address: DefaultTMP117I2CAddress,
		bus:     bus,
	}
}

/**************************************************************************************/
