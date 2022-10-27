//go:build no_runtime_type_checking

package awsglue

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

func (s *jsiiProxy_SecurityConfiguration) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_SecurityConfiguration) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateSecurityConfiguration_FromSecurityConfigurationNameParameters(scope constructs.Construct, id *string, securityConfigurationName *string) error {
	return nil
}

func validateSecurityConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecurityConfiguration_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewSecurityConfigurationParameters(scope constructs.Construct, id *string, props *SecurityConfigurationProps) error {
	return nil
}

