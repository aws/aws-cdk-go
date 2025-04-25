//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PublicKey) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PublicKey) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PublicKey) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePublicKey_FromPublicKeyIdParameters(scope constructs.Construct, id *string, publicKeyId *string) error {
	return nil
}

func validatePublicKey_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePublicKey_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePublicKey_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPublicKeyParameters(scope constructs.Construct, id *string, props *PublicKeyProps) error {
	return nil
}

