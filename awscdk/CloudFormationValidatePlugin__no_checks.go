//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationValidatePlugin) validateValidateParameters(context IPolicyValidationContext) error {
	return nil
}

func validateNewCloudFormationValidatePluginParameters(props *CloudFormationValidatePluginProps) error {
	return nil
}

