//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AgentAlias) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AgentAlias) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AgentAlias) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_AgentAlias) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AgentAlias) validateGrantGetParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AgentAlias) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AgentAlias) validateOnCloudTrailEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateAgentAlias_FromAttributesParameters(scope constructs.Construct, id *string, attrs *AgentAliasAttributes) error {
	return nil
}

func validateAgentAlias_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAgentAlias_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAgentAlias_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAgentAliasParameters(scope constructs.Construct, id *string, props *AgentAliasProps) error {
	return nil
}

