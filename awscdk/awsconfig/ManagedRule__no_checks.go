//go:build no_runtime_type_checking

package awsconfig

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedRule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (m *jsiiProxy_ManagedRule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (m *jsiiProxy_ManagedRule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (m *jsiiProxy_ManagedRule) validateOnComplianceChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (m *jsiiProxy_ManagedRule) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (m *jsiiProxy_ManagedRule) validateOnReEvaluationStatusParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateManagedRule_FromConfigRuleNameParameters(scope constructs.Construct, id *string, configRuleName *string) error {
	return nil
}

func validateManagedRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateManagedRule_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateManagedRule_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewManagedRuleParameters(scope constructs.Construct, id *string, props *ManagedRuleProps) error {
	return nil
}

