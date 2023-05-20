//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SecurityConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SecurityConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSecurityConfiguration_FromSecurityConfigurationNameParameters(scope constructs.Construct, id *string, securityConfigurationName *string) error {
	return nil
}

func validateSecurityConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecurityConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSecurityConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSecurityConfigurationParameters(scope constructs.Construct, id *string, props *SecurityConfigurationProps) error {
	return nil
}

