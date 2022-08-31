//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	"fmt"
)

func (a *jsiiProxy_Aspects) validateAddParameters(aspect IAspect) error {
	if aspect == nil {
		return fmt.Errorf("parameter aspect is required, but nil was provided")
	}

	return nil
}

func validateAspects_OfParameters(scope IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

