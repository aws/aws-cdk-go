//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_ConstructSelector) validateSelectParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateConstructSelector_ByIdParameters(pattern interface{}) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateConstructSelector_ResourcesOfTypeParameters(type_ interface{}) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

