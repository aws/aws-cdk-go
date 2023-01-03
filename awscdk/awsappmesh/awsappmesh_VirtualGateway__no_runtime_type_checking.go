//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualGateway) validateAddGatewayRouteParameters(id *string, props *GatewayRouteBaseProps) error {
	return nil
}

func (v *jsiiProxy_VirtualGateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VirtualGateway) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VirtualGateway) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VirtualGateway) validateGrantStreamAggregatedResourcesParameters(identity awsiam.IGrantable) error {
	return nil
}

func validateVirtualGateway_FromVirtualGatewayArnParameters(scope constructs.Construct, id *string, virtualGatewayArn *string) error {
	return nil
}

func validateVirtualGateway_FromVirtualGatewayAttributesParameters(scope constructs.Construct, id *string, attrs *VirtualGatewayAttributes) error {
	return nil
}

func validateVirtualGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVirtualGateway_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVirtualGateway_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVirtualGatewayParameters(scope constructs.Construct, id *string, props *VirtualGatewayProps) error {
	return nil
}

