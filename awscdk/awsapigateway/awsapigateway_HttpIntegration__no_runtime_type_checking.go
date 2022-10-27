//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpIntegration) validateBindParameters(_method Method) error {
	return nil
}

func validateNewHttpIntegrationParameters(url *string, props *HttpIntegrationProps) error {
	return nil
}

