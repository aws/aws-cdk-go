//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiKeyGrants) validateReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_ApiKeyGrants) validateReadWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_ApiKeyGrants) validateWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateApiKeyGrants_FromApiKeyParameters(resource interfacesawsapigateway.IApiKeyRef) error {
	return nil
}

