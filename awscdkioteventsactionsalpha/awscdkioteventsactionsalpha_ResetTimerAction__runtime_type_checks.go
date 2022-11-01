//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	"fmt"
)

func validateNewResetTimerActionParameters(timerName *string) error {
	if timerName == nil {
		return fmt.Errorf("parameter timerName is required, but nil was provided")
	}

	return nil
}

