//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IResource) validateAddCorsPreflightParameters(options *CorsOptions) error {
	return nil
}

func (i *jsiiProxy_IResource) validateAddMethodParameters(httpMethod *string, options *MethodOptions) error {
	return nil
}

func (i *jsiiProxy_IResource) validateAddProxyParameters(options *ProxyResourceOptions) error {
	return nil
}

func (i *jsiiProxy_IResource) validateAddResourceParameters(pathPart *string, options *ResourceOptions) error {
	return nil
}

func (i *jsiiProxy_IResource) validateGetResourceParameters(pathPart *string) error {
	return nil
}

func (i *jsiiProxy_IResource) validateResourceForPathParameters(path *string) error {
	return nil
}

