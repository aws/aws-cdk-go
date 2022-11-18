//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KeyGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KeyGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKeyGroup_FromKeyGroupIdParameters(scope constructs.Construct, id *string, keyGroupId *string) error {
	return nil
}

func validateKeyGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKeyGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKeyGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKeyGroupParameters(scope constructs.Construct, id *string, props *KeyGroupProps) error {
	return nil
}

