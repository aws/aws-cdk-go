//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeSigningConfig) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CodeSigningConfig) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CodeSigningConfig) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCodeSigningConfig_FromCodeSigningConfigArnParameters(scope constructs.Construct, id *string, codeSigningConfigArn *string) error {
	return nil
}

func validateCodeSigningConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCodeSigningConfig_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCodeSigningConfig_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCodeSigningConfigParameters(scope constructs.Construct, id *string, props *CodeSigningConfigProps) error {
	return nil
}

