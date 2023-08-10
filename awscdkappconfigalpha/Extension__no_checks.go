//go:build no_runtime_type_checking

package awscdkappconfigalpha

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Extension) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_Extension) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateExtension_FromExtensionArnParameters(scope constructs.Construct, id *string, extensionArn *string) error {
	return nil
}

func validateExtension_FromExtensionAttributesParameters(scope constructs.Construct, id *string, attr *ExtensionAttributes) error {
	return nil
}

func validateExtension_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExtension_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateExtension_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewExtensionParameters(scope constructs.Construct, id *string, props *ExtensionProps) error {
	return nil
}

