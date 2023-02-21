//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HealthCheck) validateBindParameters(scope constructs.Construct, options *HealthCheckBindOptions) error {
	return nil
}

func validateHealthCheck_GrpcParameters(options *GrpcHealthCheckOptions) error {
	return nil
}

func validateHealthCheck_HttpParameters(options *HttpHealthCheckOptions) error {
	return nil
}

func validateHealthCheck_Http2Parameters(options *HttpHealthCheckOptions) error {
	return nil
}

func validateHealthCheck_TcpParameters(options *TcpHealthCheckOptions) error {
	return nil
}

