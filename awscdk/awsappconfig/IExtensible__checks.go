//go:build !no_runtime_type_checking

package awsappconfig

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IExtensible) validateAddExtensionParameters(extension IExtension) error {
	if extension == nil {
		return fmt.Errorf("parameter extension is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IExtensible) validateAtDeploymentTickParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IExtensible) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
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

func (i *jsiiProxy_IExtensible) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IExtensible) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IExtensible) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IExtensible) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IExtensible) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IExtensible) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IExtensible) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

