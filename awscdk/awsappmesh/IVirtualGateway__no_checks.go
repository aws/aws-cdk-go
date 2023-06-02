//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVirtualGateway) validateAddGatewayRouteParameters(id *string, route *GatewayRouteBaseProps) error {
	return nil
}

func (i *jsiiProxy_IVirtualGateway) validateGrantStreamAggregatedResourcesParameters(identity awsiam.IGrantable) error {
	return nil
}

