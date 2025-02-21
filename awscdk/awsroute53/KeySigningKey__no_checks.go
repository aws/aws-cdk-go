//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeySigningKey) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KeySigningKey) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KeySigningKey) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKeySigningKey_FromKeySigningKeyAttributesParameters(scope constructs.Construct, id *string, attrs *KeySigningKeyAttributes) error {
	return nil
}

func validateKeySigningKey_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKeySigningKey_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKeySigningKey_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKeySigningKeyParameters(scope constructs.Construct, id *string, props *KeySigningKeyProps) error {
	return nil
}

