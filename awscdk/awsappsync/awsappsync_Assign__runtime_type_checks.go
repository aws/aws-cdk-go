//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func (a *jsiiProxy_Assign) validatePutInMapParameters(map_ *string) error {
	if map_ == nil {
		return fmt.Errorf("parameter map_ is required, but nil was provided")
	}

	return nil
}

func validateNewAssignParameters(attr *string, arg *string) error {
	if attr == nil {
		return fmt.Errorf("parameter attr is required, but nil was provided")
	}

	if arg == nil {
		return fmt.Errorf("parameter arg is required, but nil was provided")
	}

	return nil
}

