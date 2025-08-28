//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationInit) validateAddConfigParameters(configName *string, config InitConfig) error {
	return nil
}

func (c *jsiiProxy_CloudFormationInit) validateAddConfigSetParameters(configSetName *string) error {
	return nil
}

func (c *jsiiProxy_CloudFormationInit) validateAttachParameters(attachedResource awscdk.CfnResource, attachOptions *AttachInitOptions) error {
	return nil
}

func validateCloudFormationInit_FromConfigParameters(config InitConfig) error {
	return nil
}

func validateCloudFormationInit_FromConfigSetsParameters(props *ConfigSetProps) error {
	return nil
}

