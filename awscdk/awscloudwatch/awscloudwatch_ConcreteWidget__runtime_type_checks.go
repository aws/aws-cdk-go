//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"
)

func (c *jsiiProxy_ConcreteWidget) validatePositionParameters(x *float64, y *float64) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	if y == nil {
		return fmt.Errorf("parameter y is required, but nil was provided")
	}

	return nil
}

func validateNewConcreteWidgetParameters(width *float64, height *float64) error {
	if width == nil {
		return fmt.Errorf("parameter width is required, but nil was provided")
	}

	if height == nil {
		return fmt.Errorf("parameter height is required, but nil was provided")
	}

	return nil
}

