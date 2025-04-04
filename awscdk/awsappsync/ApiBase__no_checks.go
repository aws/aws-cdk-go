//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_ApiBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_ApiBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateApiBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApiBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateApiBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewApiBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

