//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcEndpointService) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VpcEndpointService) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VpcEndpointService) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VpcEndpointService) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (v *jsiiProxy_VpcEndpointService) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateVpcEndpointService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpcEndpointService_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewVpcEndpointServiceParameters(scope constructs.Construct, id *string, props *VpcEndpointServiceProps) error {
	return nil
}

