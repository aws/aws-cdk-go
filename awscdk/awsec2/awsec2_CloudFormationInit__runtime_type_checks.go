//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (c *jsiiProxy_CloudFormationInit) validateAddConfigParameters(configName *string, config InitConfig) error {
	if configName == nil {
		return fmt.Errorf("parameter configName is required, but nil was provided")
	}

	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudFormationInit) validateAddConfigSetParameters(configSetName *string) error {
	if configSetName == nil {
		return fmt.Errorf("parameter configSetName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudFormationInit) validateAttachParameters(attachedResource awscdk.CfnResource, attachOptions *AttachInitOptions) error {
	if attachedResource == nil {
		return fmt.Errorf("parameter attachedResource is required, but nil was provided")
	}

	if attachOptions == nil {
		return fmt.Errorf("parameter attachOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attachOptions, func() string { return "parameter attachOptions" }); err != nil {
		return err
	}

	return nil
}

func validateCloudFormationInit_FromConfigParameters(config InitConfig) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}

	return nil
}

func validateCloudFormationInit_FromConfigSetsParameters(props *ConfigSetProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

