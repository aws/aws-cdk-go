//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MockIntegration) validateBindParameters(_method Method) error {
	return nil
}

func validateNewMockIntegrationParameters(options *IntegrationOptions) error {
	return nil
}

