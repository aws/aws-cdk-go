//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpGatewayRoutePathMatch) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateHttpGatewayRoutePathMatch_ExactlyParameters(path *string) error {
	return nil
}

func validateHttpGatewayRoutePathMatch_RegexParameters(regex *string) error {
	return nil
}

func validateHttpGatewayRoutePathMatch_StartsWithParameters(prefix *string) error {
	return nil
}

