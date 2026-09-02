//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_Network) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (n *jsiiProxy_Network) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_Network) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_Network) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNetwork_FromNetworkAttributesParameters(scope constructs.Construct, id *string, attrs *NetworkAttributes) error {
	return nil
}

func validateNetwork_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNetwork_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNetwork_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNetworkParameters(scope constructs.Construct, id *string, props *NetworkProps) error {
	return nil
}

