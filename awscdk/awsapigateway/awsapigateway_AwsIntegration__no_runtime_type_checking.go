//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsIntegration) validateBindParameters(method Method) error {
	return nil
}

func validateNewAwsIntegrationParameters(props *AwsIntegrationProps) error {
	return nil
}

