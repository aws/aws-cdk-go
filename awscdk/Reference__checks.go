//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (r *jsiiProxy_Reference) validateNewErrorParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_Reference) validateResolveParameters(_context IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func validateReference_IsReferenceParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewReferenceParameters(value interface{}, target constructs.IConstruct) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

