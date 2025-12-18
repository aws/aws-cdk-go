//go:build !no_runtime_type_checking

package awsappconfig

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (i *jsiiProxy_IEnvironment) validateAddDeploymentParameters(configuration IConfiguration) error {
	if configuration == nil {
		return fmt.Errorf("parameter configuration is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validateAddExtensionParameters(extension IExtension) error {
	if extension == nil {
		return fmt.Errorf("parameter extension is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validateAtDeploymentTickParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validateGrantParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validateGrantReadConfigParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
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

func (i *jsiiProxy_IEnvironment) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	if eventDestination == nil {
		return fmt.Errorf("parameter eventDestination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IEnvironment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

