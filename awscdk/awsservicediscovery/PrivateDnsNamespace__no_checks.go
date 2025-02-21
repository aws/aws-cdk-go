//go:build no_runtime_type_checking

package awsservicediscovery

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateDnsNamespace) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PrivateDnsNamespace) validateCreateServiceParameters(id *string, props *DnsServiceProps) error {
	return nil
}

func (p *jsiiProxy_PrivateDnsNamespace) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PrivateDnsNamespace) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePrivateDnsNamespace_FromPrivateDnsNamespaceAttributesParameters(scope constructs.Construct, id *string, attrs *PrivateDnsNamespaceAttributes) error {
	return nil
}

func validatePrivateDnsNamespace_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePrivateDnsNamespace_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePrivateDnsNamespace_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPrivateDnsNamespaceParameters(scope constructs.Construct, id *string, props *PrivateDnsNamespaceProps) error {
	return nil
}

