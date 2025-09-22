//go:build !no_runtime_type_checking

package awscdkioteventsactionsalpha

import (
	"fmt"
)

func validateNewClearTimerActionParameters(timerName *string) error {
	if timerName == nil {
		return fmt.Errorf("parameter timerName is required, but nil was provided")
	}

	return nil
}

