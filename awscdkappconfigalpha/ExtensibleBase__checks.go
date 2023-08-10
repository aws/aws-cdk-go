//go:build !no_runtime_type_checking

package awscdkappconfigalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_ExtensibleBase) validateAddExtensionParameters(extension IExtension) error {
	if extension == nil {
		return fmt.Errorf("parameter extension is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
	if actionPoint == "" {
		return fmt.Errorf("parameter actionPoint is required, but nil was provided")
	}

	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExtensibleBase) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExtensibleBase) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewExtensibleBaseParameters(scope constructs.Construct, resourceArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if resourceArn == nil {
		return fmt.Errorf("parameter resourceArn is required, but nil was provided")
	}

	return nil
}

