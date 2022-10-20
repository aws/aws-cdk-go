//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsglobalaccelerator

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointGroup) validateAddEndpointParameters(endpoint IEndpoint) error {
	return nil
}

func (e *jsiiProxy_EndpointGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EndpointGroup) validateConnectionsPeerParameters(id *string, vpc awsec2.IVpc) error {
	return nil
}

func (e *jsiiProxy_EndpointGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EndpointGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_EndpointGroup) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (e *jsiiProxy_EndpointGroup) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateEndpointGroup_FromEndpointGroupArnParameters(scope constructs.Construct, id *string, endpointGroupArn *string) error {
	return nil
}

func validateEndpointGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEndpointGroup_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewEndpointGroupParameters(scope constructs.Construct, id *string, props *EndpointGroupProps) error {
	return nil
}

