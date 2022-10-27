//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"
)

func (i *jsiiProxy_IWidget) validatePositionParameters(x *float64, y *float64) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	if y == nil {
		return fmt.Errorf("parameter y is required, but nil was provided")
	}

	return nil
}

