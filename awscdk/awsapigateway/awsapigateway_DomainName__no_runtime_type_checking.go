//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DomainName) validateAddApiMappingParameters(targetStage IStage, options *ApiMappingOptions) error {
	return nil
}

func (d *jsiiProxy_DomainName) validateAddBasePathMappingParameters(targetApi IRestApi, options *BasePathMappingOptions) error {
	return nil
}

func (d *jsiiProxy_DomainName) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DomainName) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DomainName) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDomainName_FromDomainNameAttributesParameters(scope constructs.Construct, id *string, attrs *DomainNameAttributes) error {
	return nil
}

func validateDomainName_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDomainName_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDomainName_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDomainNameParameters(scope constructs.Construct, id *string, props *DomainNameProps) error {
	return nil
}

