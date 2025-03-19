//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayRouteTableAssociation) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTableAssociation) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTableAssociation) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGatewayRouteTableAssociation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayRouteTableAssociation_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayRouteTableAssociation_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayRouteTableAssociationParameters(scope constructs.Construct, id *string, props *TransitGatewayRouteTableAssociationProps) error {
	return nil
}

