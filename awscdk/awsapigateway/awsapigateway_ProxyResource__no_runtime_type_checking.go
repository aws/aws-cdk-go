//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProxyResource) validateAddCorsPreflightParameters(options *CorsOptions) error {
	return nil
}

func (p *jsiiProxy_ProxyResource) validateAddMethodParameters(httpMethod *string, options *MethodOptions) error {
	return nil
}

func (p *jsiiProxy_ProxyResource) validateAddProxyParameters(options *ProxyResourceOptions) error {
	return nil
}

func (p *jsiiProxy_ProxyResource) validateAddResourceParameters(pathPart *string, options *ResourceOptions) error {
	return nil
}

func (p *jsiiProxy_ProxyResource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_ProxyResource) validateGetResourceParameters(pathPart *string) error {
	return nil
}

func (p *jsiiProxy_ProxyResource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_ProxyResource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_ProxyResource) validateResourceForPathParameters(path *string) error {
	return nil
}

func validateProxyResource_FromResourceAttributesParameters(scope constructs.Construct, id *string, attrs *ResourceAttributes) error {
	return nil
}

func validateProxyResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProxyResource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateProxyResource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewProxyResourceParameters(scope constructs.Construct, id *string, props *ProxyResourceProps) error {
	return nil
}

