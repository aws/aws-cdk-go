//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkConnector) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (n *jsiiProxy_NetworkConnector) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NetworkConnector) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NetworkConnector) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNetworkConnector_FromNetworkConnectorArnParameters(scope constructs.Construct, id *string, networkConnectorArn *string) error {
	return nil
}

func validateNetworkConnector_FromNetworkConnectorAttributesParameters(scope constructs.Construct, id *string, attrs *NetworkConnectorAttributes) error {
	return nil
}

func validateNetworkConnector_FromNetworkConnectorVersionArnParameters(scope constructs.Construct, id *string, networkConnectorVersionArn *string) error {
	return nil
}

func validateNetworkConnector_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNetworkConnector_IsNetworkConnectorParameters(x interface{}) error {
	return nil
}

func validateNetworkConnector_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNetworkConnector_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNetworkConnectorParameters(scope constructs.Construct, id *string, props *NetworkConnectorProps) error {
	return nil
}

