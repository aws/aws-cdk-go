//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SamlProvider) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SamlProvider) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SamlProvider) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSamlProvider_FromSamlProviderArnParameters(scope constructs.Construct, id *string, samlProviderArn *string) error {
	return nil
}

func validateSamlProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSamlProvider_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSamlProvider_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSamlProviderParameters(scope constructs.Construct, id *string, props *SamlProviderProps) error {
	return nil
}

