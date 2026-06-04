//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CrossStackReferences) validateConsumeParameters(strength ReferenceStrength) error {
	if strength == "" {
		return fmt.Errorf("parameter strength is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CrossStackReferences) validateProduceParameters(strength ReferenceStrength) error {
	if strength == "" {
		return fmt.Errorf("parameter strength is required, but nil was provided")
	}

	return nil
}

func validateCrossStackReferences_OfParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

