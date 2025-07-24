//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_PropertyInjectors) validateForParameters(uniqueId *string) error {
	if uniqueId == nil {
		return fmt.Errorf("parameter uniqueId is required, but nil was provided")
	}

	return nil
}

func validatePropertyInjectors_HasPropertyInjectorsParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validatePropertyInjectors_OfParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

