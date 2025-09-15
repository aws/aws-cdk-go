//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AgentBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AgentBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AgentBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_AgentBase) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AgentBase) validateMetricCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_AgentBase) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateAgentBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAgentBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAgentBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAgentBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

