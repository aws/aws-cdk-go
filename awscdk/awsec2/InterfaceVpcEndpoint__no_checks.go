//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InterfaceVpcEndpoint) validateAddToPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_InterfaceVpcEndpoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_InterfaceVpcEndpoint) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_InterfaceVpcEndpoint) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateInterfaceVpcEndpoint_FromInterfaceVpcEndpointAttributesParameters(scope constructs.Construct, id *string, attrs *InterfaceVpcEndpointAttributes) error {
	return nil
}

func validateInterfaceVpcEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInterfaceVpcEndpoint_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateInterfaceVpcEndpoint_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewInterfaceVpcEndpointParameters(scope constructs.Construct, id *string, props *InterfaceVpcEndpointProps) error {
	return nil
}

