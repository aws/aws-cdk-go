//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"
)

func (d *jsiiProxy_DynamoProjectionExpression) validateAtIndexParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DynamoProjectionExpression) validateWithAttributeParameters(attr *string) error {
	if attr == nil {
		return fmt.Errorf("parameter attr is required, but nil was provided")
	}

	return nil
}

