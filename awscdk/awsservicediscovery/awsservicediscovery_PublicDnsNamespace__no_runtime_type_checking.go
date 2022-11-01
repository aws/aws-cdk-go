//go:build no_runtime_type_checking

package awsservicediscovery

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PublicDnsNamespace) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PublicDnsNamespace) validateCreateServiceParameters(id *string, props *DnsServiceProps) error {
	return nil
}

func (p *jsiiProxy_PublicDnsNamespace) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PublicDnsNamespace) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PublicDnsNamespace) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (p *jsiiProxy_PublicDnsNamespace) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validatePublicDnsNamespace_FromPublicDnsNamespaceAttributesParameters(scope constructs.Construct, id *string, attrs *PublicDnsNamespaceAttributes) error {
	return nil
}

func validatePublicDnsNamespace_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePublicDnsNamespace_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewPublicDnsNamespaceParameters(scope constructs.Construct, id *string, props *PublicDnsNamespaceProps) error {
	return nil
}

