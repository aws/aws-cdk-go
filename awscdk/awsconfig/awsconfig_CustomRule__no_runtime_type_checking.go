//go:build no_runtime_type_checking

package awsconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomRule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CustomRule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CustomRule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CustomRule) validateOnComplianceChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (c *jsiiProxy_CustomRule) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (c *jsiiProxy_CustomRule) validateOnReEvaluationStatusParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateCustomRule_FromConfigRuleNameParameters(scope constructs.Construct, id *string, configRuleName *string) error {
	return nil
}

func validateCustomRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomRule_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCustomRule_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCustomRuleParameters(scope constructs.Construct, id *string, props *CustomRuleProps) error {
	return nil
}

