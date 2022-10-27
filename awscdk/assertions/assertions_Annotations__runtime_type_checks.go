//go:build !no_runtime_type_checking

package assertions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (a *jsiiProxy_Annotations) validateFindErrorParameters(constructPath *string, message interface{}) error {
	if constructPath == nil {
		return fmt.Errorf("parameter constructPath is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateFindInfoParameters(constructPath *string, message interface{}) error {
	if constructPath == nil {
		return fmt.Errorf("parameter constructPath is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateFindWarningParameters(constructPath *string, message interface{}) error {
	if constructPath == nil {
		return fmt.Errorf("parameter constructPath is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateHasErrorParameters(constructPath *string, message interface{}) error {
	if constructPath == nil {
		return fmt.Errorf("parameter constructPath is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateHasInfoParameters(constructPath *string, message interface{}) error {
	if constructPath == nil {
		return fmt.Errorf("parameter constructPath is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateHasNoErrorParameters(constructPath *string, message interface{}) error {
	if constructPath == nil {
		return fmt.Errorf("parameter constructPath is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateHasNoInfoParameters(constructPath *string, message interface{}) error {
	if constructPath == nil {
		return fmt.Errorf("parameter constructPath is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateHasNoWarningParameters(constructPath *string, message interface{}) error {
	if constructPath == nil {
		return fmt.Errorf("parameter constructPath is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateHasWarningParameters(constructPath *string, message interface{}) error {
	if constructPath == nil {
		return fmt.Errorf("parameter constructPath is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func validateAnnotations_FromStackParameters(stack awscdk.Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

