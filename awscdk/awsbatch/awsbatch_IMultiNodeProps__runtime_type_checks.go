//go:build !no_runtime_type_checking

package awsbatch

import (
	"fmt"
)

func (j *jsiiProxy_IMultiNodeProps) validateSetCountParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IMultiNodeProps) validateSetMainNodeParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IMultiNodeProps) validateSetRangePropsParameters(val *[]INodeRangeProps) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

