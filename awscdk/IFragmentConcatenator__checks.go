//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (i *jsiiProxy_IFragmentConcatenator) validateJoinParameters(left interface{}, right interface{}) error {
	if left == nil {
		return fmt.Errorf("parameter left is required, but nil was provided")
	}

	if right == nil {
		return fmt.Errorf("parameter right is required, but nil was provided")
	}

	return nil
}

