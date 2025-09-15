//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Agent) validateAddActionGroupParameters(actionGroup AgentActionGroup) error {
	return nil
}

func (a *jsiiProxy_Agent) validateAddGuardrailParameters(guardrail IGuardrail) error {
	return nil
}

func (a *jsiiProxy_Agent) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Agent) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Agent) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_Agent) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_Agent) validateMetricCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Agent) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateAgent_FromAgentAttributesParameters(scope constructs.Construct, id *string, attrs *AgentAttributes) error {
	return nil
}

func validateAgent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAgent_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAgent_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAgentParameters(scope constructs.Construct, id *string, props *AgentProps) error {
	return nil
}

