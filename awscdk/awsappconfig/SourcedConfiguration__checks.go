//go:build !no_runtime_type_checking

package awsappconfig

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SourcedConfiguration) validateAddExtensionParameters(extension IExtension) error {
	if extension == nil {
		return fmt.Errorf("parameter extension is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SourcedConfiguration) validateAtDeploymentTickParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SourcedConfiguration) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
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

func (s *jsiiProxy_SourcedConfiguration) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SourcedConfiguration) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SourcedConfiguration) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SourcedConfiguration) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SourcedConfiguration) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SourcedConfiguration) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SourcedConfiguration) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSourcedConfiguration_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SourcedConfiguration) validateSetApplicationIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SourcedConfiguration) validateSetExtensibleParameters(val ExtensibleBase) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSourcedConfigurationParameters(scope constructs.Construct, id *string, props *SourcedConfigurationProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

