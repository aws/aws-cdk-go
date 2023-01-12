//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouteSpec) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateRouteSpec_GrpcParameters(options *GrpcRouteSpecOptions) error {
	return nil
}

func validateRouteSpec_HttpParameters(options *HttpRouteSpecOptions) error {
	return nil
}

func validateRouteSpec_Http2Parameters(options *HttpRouteSpecOptions) error {
	return nil
}

func validateRouteSpec_TcpParameters(options *TcpRouteSpecOptions) error {
	return nil
}

