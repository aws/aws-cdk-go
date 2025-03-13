//go:build no_runtime_type_checking

package awsconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CustomPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CustomPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CustomPolicy) validateOnComplianceChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (c *jsiiProxy_CustomPolicy) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (c *jsiiProxy_CustomPolicy) validateOnReEvaluationStatusParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateCustomPolicy_FromConfigRuleNameParameters(scope constructs.Construct, id *string, configRuleName *string) error {
	return nil
}

func validateCustomPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCustomPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCustomPolicyParameters(scope constructs.Construct, id *string, props *CustomPolicyProps) error {
	return nil
}

