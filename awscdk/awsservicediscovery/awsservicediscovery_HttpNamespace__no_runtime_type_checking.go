//go:build no_runtime_type_checking

package awsservicediscovery

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpNamespace) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (h *jsiiProxy_HttpNamespace) validateCreateServiceParameters(id *string, props *BaseServiceProps) error {
	return nil
}

func (h *jsiiProxy_HttpNamespace) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (h *jsiiProxy_HttpNamespace) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateHttpNamespace_FromHttpNamespaceAttributesParameters(scope constructs.Construct, id *string, attrs *HttpNamespaceAttributes) error {
	return nil
}

func validateHttpNamespace_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHttpNamespace_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateHttpNamespace_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewHttpNamespaceParameters(scope constructs.Construct, id *string, props *HttpNamespaceProps) error {
	return nil
}

