//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_Annotations) validateAcknowledgeWarningParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateAddDeprecationParameters(api *string, message *string) error {
	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateAddErrorParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateAddInfoParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateAddWarningParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Annotations) validateAddWarningV2Parameters(id *string, message *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func validateAnnotations_OfParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

