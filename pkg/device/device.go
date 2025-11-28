/**************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@license	Copyright © 2021-2026 Michael Roberts

/**************************************************************************************/

package device

/**************************************************************************************/

import (
	"machine"
)

/**************************************************************************************/

// Device represents a generic device connected over I2C.
type Device struct {
	Address uint16      // I2C address of the sensor
	Bus     machine.I2C // I2C bus to which the sensor is connected
}

/**************************************************************************************/
