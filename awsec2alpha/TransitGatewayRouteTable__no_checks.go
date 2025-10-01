//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayRouteTable) validateAddAssociationParameters(id *string, transitGatewayAttachment ITransitGatewayAttachment) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTable) validateAddBlackholeRouteParameters(id *string, destinationCidr *string) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTable) validateAddRouteParameters(id *string, transitGatewayAttachment ITransitGatewayAttachment, destinationCidr *string) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTable) validateEnablePropagationParameters(id *string, transitGatewayAttachment ITransitGatewayAttachment) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTable) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayRouteTable) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTransitGatewayRouteTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayRouteTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayRouteTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayRouteTableParameters(scope constructs.Construct, id *string, props *TransitGatewayRouteTableProps) error {
	return nil
}

