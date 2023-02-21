//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ParameterGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_ParameterGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_ParameterGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateParameterGroup_FromParameterGroupNameParameters(scope constructs.Construct, id *string, parameterGroupName *string) error {
	return nil
}

func validateParameterGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateParameterGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateParameterGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewParameterGroupParameters(scope constructs.Construct, id *string, props *ParameterGroupProps) error {
	return nil
}

