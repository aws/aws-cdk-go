//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Integration) validateBindParameters(_method Method) error {
	return nil
}

func validateNewIntegrationParameters(props *IntegrationProps) error {
	return nil
}

