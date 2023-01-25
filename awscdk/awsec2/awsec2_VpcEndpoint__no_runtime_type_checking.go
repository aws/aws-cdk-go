//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcEndpoint) validateAddToPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (v *jsiiProxy_VpcEndpoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VpcEndpoint) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VpcEndpoint) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VpcEndpoint) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (v *jsiiProxy_VpcEndpoint) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateVpcEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpcEndpoint_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewVpcEndpointParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

