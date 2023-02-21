//go:build no_runtime_type_checking

package awsconfig

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessKeysRotated) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AccessKeysRotated) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AccessKeysRotated) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_AccessKeysRotated) validateOnComplianceChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (a *jsiiProxy_AccessKeysRotated) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (a *jsiiProxy_AccessKeysRotated) validateOnReEvaluationStatusParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateAccessKeysRotated_FromConfigRuleNameParameters(scope constructs.Construct, id *string, configRuleName *string) error {
	return nil
}

func validateAccessKeysRotated_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAccessKeysRotated_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAccessKeysRotated_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAccessKeysRotatedParameters(scope constructs.Construct, id *string, props *AccessKeysRotatedProps) error {
	return nil
}

