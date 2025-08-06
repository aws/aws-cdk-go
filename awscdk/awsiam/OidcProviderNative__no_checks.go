//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OidcProviderNative) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (o *jsiiProxy_OidcProviderNative) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (o *jsiiProxy_OidcProviderNative) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateOidcProviderNative_FromOidcProviderArnParameters(scope constructs.Construct, id *string, oidcProviderArn *string) error {
	return nil
}

func validateOidcProviderNative_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOidcProviderNative_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateOidcProviderNative_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewOidcProviderNativeParameters(scope constructs.Construct, id *string, props *OidcProviderNativeProps) error {
	return nil
}

