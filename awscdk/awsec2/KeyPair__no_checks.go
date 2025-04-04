//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyPair) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KeyPair) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKeyPair_FromKeyPairAttributesParameters(scope constructs.Construct, id *string, attrs *KeyPairAttributes) error {
	return nil
}

func validateKeyPair_FromKeyPairNameParameters(scope constructs.Construct, id *string, keyPairName *string) error {
	return nil
}

func validateKeyPair_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKeyPair_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKeyPair_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKeyPairParameters(scope constructs.Construct, id *string, props *KeyPairProps) error {
	return nil
}

