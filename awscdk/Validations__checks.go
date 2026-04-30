//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (v *jsiiProxy_Validations) validateAcknowledgeParameters(rules *[]*Acknowledgment) error {
	for idx_6c621d, v := range *rules {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter rules[%#v]", idx_6c621d) }); err != nil {
			return err
		}
	}

	return nil
}

func (v *jsiiProxy_Validations) validateAddErrorParameters(id *string, message *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_Validations) validateAddWarningParameters(id *string, message *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func validateValidations_OfParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

