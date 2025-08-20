//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AgentAliasBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AgentAliasBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AgentAliasBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_AgentAliasBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AgentAliasBase) validateGrantGetParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AgentAliasBase) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_AgentAliasBase) validateOnCloudTrailEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateAgentAliasBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAgentAliasBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAgentAliasBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAgentAliasBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

