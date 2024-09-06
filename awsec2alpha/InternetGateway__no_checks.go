//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InternetGateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_InternetGateway) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_InternetGateway) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateInternetGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInternetGateway_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateInternetGateway_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewInternetGatewayParameters(scope constructs.Construct, id *string, props *InternetGatewayProps) error {
	return nil
}

