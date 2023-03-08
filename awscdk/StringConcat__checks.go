//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"
)

func (s *jsiiProxy_StringConcat) validateJoinParameters(left interface{}, right interface{}) error {
	if left == nil {
		return fmt.Errorf("parameter left is required, but nil was provided")
	}

	if right == nil {
		return fmt.Errorf("parameter right is required, but nil was provided")
	}

	return nil
}

