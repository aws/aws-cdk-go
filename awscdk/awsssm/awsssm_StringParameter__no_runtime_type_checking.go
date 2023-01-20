//go:build no_runtime_type_checking

package awsssm

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StringParameter) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_StringParameter) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_StringParameter) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_StringParameter) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StringParameter) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateStringParameter_FromSecureStringParameterAttributesParameters(scope constructs.Construct, id *string, attrs *SecureStringParameterAttributes) error {
	return nil
}

func validateStringParameter_FromStringParameterAttributesParameters(scope constructs.Construct, id *string, attrs *StringParameterAttributes) error {
	return nil
}

func validateStringParameter_FromStringParameterNameParameters(scope constructs.Construct, id *string, stringParameterName *string) error {
	return nil
}

func validateStringParameter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStringParameter_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateStringParameter_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateStringParameter_ValueForSecureStringParameterParameters(scope constructs.Construct, parameterName *string, version *float64) error {
	return nil
}

func validateStringParameter_ValueForStringParameterParameters(scope constructs.Construct, parameterName *string) error {
	return nil
}

func validateStringParameter_ValueForTypedStringParameterParameters(scope constructs.Construct, parameterName *string) error {
	return nil
}

func validateStringParameter_ValueForTypedStringParameterV2Parameters(scope constructs.Construct, parameterName *string) error {
	return nil
}

func validateStringParameter_ValueFromLookupParameters(scope constructs.Construct, parameterName *string) error {
	return nil
}

func validateNewStringParameterParameters(scope constructs.Construct, id *string, props *StringParameterProps) error {
	return nil
}

