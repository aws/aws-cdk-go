//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGateway) validateAddRouteTableParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateAttachVpcParameters(id *string, options *AttachVpcOptions) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGateway) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGateway_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGateway_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayParameters(scope constructs.Construct, id *string, props *TransitGatewayProps) error {
	return nil
}

