//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func (i *jsiiProxy_IEventSourceDlq) validateBindParameters(target IEventSourceMapping, targetHandler IFunction) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if targetHandler == nil {
		return fmt.Errorf("parameter targetHandler is required, but nil was provided")
	}

	return nil
}

