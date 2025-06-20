//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SagemakerIntegration) validateBindParameters(method Method) error {
	return nil
}

func validateNewSagemakerIntegrationParameters(endpoint awssagemaker.IEndpoint, options *SagemakerIntegrationOptions) error {
	return nil
}

