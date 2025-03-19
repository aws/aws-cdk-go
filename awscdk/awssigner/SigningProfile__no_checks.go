//go:build no_runtime_type_checking

package awssigner

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SigningProfile) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SigningProfile) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SigningProfile) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSigningProfile_FromSigningProfileAttributesParameters(scope constructs.Construct, id *string, attrs *SigningProfileAttributes) error {
	return nil
}

func validateSigningProfile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSigningProfile_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSigningProfile_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSigningProfileParameters(scope constructs.Construct, id *string, props *SigningProfileProps) error {
	return nil
}

