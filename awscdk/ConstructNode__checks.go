//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

func (c *jsiiProxy_ConstructNode) validateAddErrorParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructNode) validateAddInfoParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructNode) validateAddMetadataParameters(type_ *string, data interface{}) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructNode) validateAddValidationParameters(validation constructs.IValidation) error {
	if validation == nil {
		return fmt.Errorf("parameter validation is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructNode) validateAddWarningParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructNode) validateApplyAspectParameters(aspect IAspect) error {
	if aspect == nil {
		return fmt.Errorf("parameter aspect is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructNode) validateFindChildParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructNode) validateSetContextParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructNode) validateTryFindChildParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructNode) validateTryGetContextParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstructNode) validateTryRemoveChildParameters(childName *string) error {
	if childName == nil {
		return fmt.Errorf("parameter childName is required, but nil was provided")
	}

	return nil
}

func validateConstructNode_PrepareParameters(node ConstructNode) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func validateConstructNode_SynthParameters(node ConstructNode, options *SynthesisOptions) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateConstructNode_ValidateParameters(node ConstructNode) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func validateNewConstructNodeParameters(host Construct, scope IConstruct, id *string) error {
	if host == nil {
		return fmt.Errorf("parameter host is required, but nil was provided")
	}

	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

