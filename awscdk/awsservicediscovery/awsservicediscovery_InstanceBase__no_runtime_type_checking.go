//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsservicediscovery

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_InstanceBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_InstanceBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_InstanceBase) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (i *jsiiProxy_InstanceBase) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateInstanceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInstanceBase_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewInstanceBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

