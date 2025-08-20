//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GatewayRouteHostnameMatch) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateGatewayRouteHostnameMatch_EndsWithParameters(suffix *string) error {
	return nil
}

func validateGatewayRouteHostnameMatch_ExactlyParameters(name *string) error {
	return nil
}

