//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Resource) validateAddCorsPreflightParameters(options *CorsOptions) error {
	return nil
}

func (r *jsiiProxy_Resource) validateAddMethodParameters(httpMethod *string, options *MethodOptions) error {
	return nil
}

func (r *jsiiProxy_Resource) validateAddProxyParameters(options *ProxyResourceOptions) error {
	return nil
}

func (r *jsiiProxy_Resource) validateAddResourceParameters(pathPart *string, options *ResourceOptions) error {
	return nil
}

func (r *jsiiProxy_Resource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetResourceParameters(pathPart *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateResourceForPathParameters(path *string) error {
	return nil
}

func validateResource_FromResourceAttributesParameters(scope constructs.Construct, id *string, attrs *ResourceAttributes) error {
	return nil
}

func validateResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateResource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewResourceParameters(scope constructs.Construct, id *string, props *ResourceProps) error {
	return nil
}

