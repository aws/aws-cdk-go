//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"
)

func (r *jsiiProxy_Row) validateAddWidgetParameters(w IWidget) error {
	if w == nil {
		return fmt.Errorf("parameter w is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_Row) validatePositionParameters(x *float64, y *float64) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	if y == nil {
		return fmt.Errorf("parameter y is required, but nil was provided")
	}

	return nil
}

