//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BridgeSource) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (b *jsiiProxy_BridgeSource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BridgeSource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BridgeSource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateBridgeSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBridgeSource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBridgeSource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBridgeSourceParameters(scope constructs.Construct, id *string, props *BridgeSourceProps) error {
	return nil
}

