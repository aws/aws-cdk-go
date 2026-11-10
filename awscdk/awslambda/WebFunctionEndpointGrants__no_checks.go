//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebFunctionEndpointGrants) validateInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateWebFunctionEndpointGrants_FromWebFunctionEndpointParameters(endpoint interfacesawslambda.IWebFunctionEndpointRef) error {
	return nil
}

