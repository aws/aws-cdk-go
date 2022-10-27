//go:build !no_runtime_type_checking

package awsautoscalingcommon

import (
	"fmt"
)

func (i *jsiiProxy_IRandomGenerator) validateNextIntParameters(min *float64, max *float64) error {
	if min == nil {
		return fmt.Errorf("parameter min is required, but nil was provided")
	}

	if max == nil {
		return fmt.Errorf("parameter max is required, but nil was provided")
	}

	return nil
}

