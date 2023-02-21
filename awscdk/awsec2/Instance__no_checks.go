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

func validateInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInstance_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateInstance_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewInstanceParameters(scope constructs.Construct, id *string, props *InstanceProps) error {
	return nil
}

