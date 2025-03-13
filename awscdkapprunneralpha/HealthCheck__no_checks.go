//go:build no_runtime_type_checking

package awscdkapprunneralpha

// Building without runtime type checking enabled, so all the below just return nil

func validateHealthCheck_HttpParameters(options *HttpHealthCheckOptions) error {
	return nil
}

func validateHealthCheck_TcpParameters(options *TcpHealthCheckOptions) error {
	return nil
}

