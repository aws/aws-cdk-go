//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITransitGatewayRouteTable) validateAddAssociationParameters(id *string, transitGatewayAttachment ITransitGatewayAttachment) error {
	return nil
}

func (i *jsiiProxy_ITransitGatewayRouteTable) validateAddBlackholeRouteParameters(id *string, destinationCidr *string) error {
	return nil
}

func (i *jsiiProxy_ITransitGatewayRouteTable) validateAddRouteParameters(id *string, transitGatewayAttachment ITransitGatewayAttachment, destinationCidr *string) error {
	return nil
}

func (i *jsiiProxy_ITransitGatewayRouteTable) validateEnablePropagationParameters(id *string, transitGatewayAttachment ITransitGatewayAttachment) error {
	return nil
}

func (i *jsiiProxy_ITransitGatewayRouteTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

