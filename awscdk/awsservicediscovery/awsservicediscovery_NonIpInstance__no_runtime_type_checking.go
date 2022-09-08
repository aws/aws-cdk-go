//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsservicediscovery

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NonIpInstance) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NonIpInstance) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NonIpInstance) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (n *jsiiProxy_NonIpInstance) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (n *jsiiProxy_NonIpInstance) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateNonIpInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNonIpInstance_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewNonIpInstanceParameters(scope constructs.Construct, id *string, props *NonIpInstanceProps) error {
	return nil
}

