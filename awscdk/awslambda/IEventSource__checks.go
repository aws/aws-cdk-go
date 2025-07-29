//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func (i *jsiiProxy_IEventSource) validateBindParameters(target IFunction) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

