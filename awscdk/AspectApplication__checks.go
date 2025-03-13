//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (j *jsiiProxy_AspectApplication) validateSetPriorityParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAspectApplicationParameters(construct constructs.IConstruct, aspect IAspect, priority *float64) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	if aspect == nil {
		return fmt.Errorf("parameter aspect is required, but nil was provided")
	}

	if priority == nil {
		return fmt.Errorf("parameter priority is required, but nil was provided")
	}

	return nil
}

