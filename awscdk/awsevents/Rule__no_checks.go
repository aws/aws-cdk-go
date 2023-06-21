//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Rule) validateAddEventPatternParameters(eventPattern *EventPattern) error {
	return nil
}

func (r *jsiiProxy_Rule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_Rule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_Rule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRule_FromEventRuleArnParameters(scope constructs.Construct, id *string, eventRuleArn *string) error {
	return nil
}

func validateRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRule_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRule_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRuleParameters(scope constructs.Construct, id *string, props *RuleProps) error {
	return nil
}

