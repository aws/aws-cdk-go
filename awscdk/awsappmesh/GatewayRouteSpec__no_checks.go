//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GatewayRouteSpec) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateGatewayRouteSpec_GrpcParameters(options *GrpcGatewayRouteSpecOptions) error {
	return nil
}

func validateGatewayRouteSpec_HttpParameters(options *HttpGatewayRouteSpecOptions) error {
	return nil
}

func validateGatewayRouteSpec_Http2Parameters(options *HttpGatewayRouteSpecOptions) error {
	return nil
}

