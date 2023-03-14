//go:build !no_runtime_type_checking

// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	"fmt"
)

func validateNewSetTimerActionParameters(timerName *string, timerDuration TimerDuration) error {
	if timerName == nil {
		return fmt.Errorf("parameter timerName is required, but nil was provided")
	}

	if timerDuration == nil {
		return fmt.Errorf("parameter timerDuration is required, but nil was provided")
	}

	return nil
}

