//go:build no_runtime_type_checking
// +build no_runtime_type_checking

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

func (r *jsiiProxy_Rule) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (r *jsiiProxy_Rule) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateRule_FromEventRuleArnParameters(scope constructs.Construct, id *string, eventRuleArn *string) error {
	return nil
}

func validateRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRule_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewRuleParameters(scope constructs.Construct, id *string, props *RuleProps) error {
	return nil
}

