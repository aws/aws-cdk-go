//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppsyncFunction) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AppsyncFunction) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AppsyncFunction) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_AppsyncFunction) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_AppsyncFunction) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateAppsyncFunction_FromAppsyncFunctionAttributesParameters(scope constructs.Construct, id *string, attrs *AppsyncFunctionAttributes) error {
	return nil
}

func validateAppsyncFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAppsyncFunction_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewAppsyncFunctionParameters(scope constructs.Construct, id *string, props *AppsyncFunctionProps) error {
	return nil
}

