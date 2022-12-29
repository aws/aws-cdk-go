//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourceBase) validateAddCorsPreflightParameters(options *CorsOptions) error {
	return nil
}

func (r *jsiiProxy_ResourceBase) validateAddMethodParameters(httpMethod *string, options *MethodOptions) error {
	return nil
}

func (r *jsiiProxy_ResourceBase) validateAddProxyParameters(options *ProxyResourceOptions) error {
	return nil
}

func (r *jsiiProxy_ResourceBase) validateAddResourceParameters(pathPart *string, options *ResourceOptions) error {
	return nil
}

func (r *jsiiProxy_ResourceBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_ResourceBase) validateGetResourceParameters(pathPart *string) error {
	return nil
}

func (r *jsiiProxy_ResourceBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_ResourceBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_ResourceBase) validateResourceForPathParameters(path *string) error {
	return nil
}

func validateResourceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResourceBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateResourceBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewResourceBaseParameters(scope constructs.Construct, id *string) error {
	return nil
}

