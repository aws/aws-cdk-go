//go:build no_runtime_type_checking

package awsservicediscovery

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AliasTargetInstance) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AliasTargetInstance) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AliasTargetInstance) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAliasTargetInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAliasTargetInstance_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAliasTargetInstance_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAliasTargetInstanceParameters(scope constructs.Construct, id *string, props *AliasTargetInstanceProps) error {
	return nil
}

