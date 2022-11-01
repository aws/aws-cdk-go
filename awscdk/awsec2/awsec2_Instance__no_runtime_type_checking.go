//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Instance) validateAddSecurityGroupParameters(securityGroup ISecurityGroup) error {
	return nil
}

func (i *jsiiProxy_Instance) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_Instance) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_Instance) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_Instance) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_Instance) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (i *jsiiProxy_Instance) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInstance_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewInstanceParameters(scope constructs.Construct, id *string, props *InstanceProps) error {
	return nil
}

