//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeEnvironment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ComputeEnvironment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ComputeEnvironment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_ComputeEnvironment) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_ComputeEnvironment) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateComputeEnvironment_FromComputeEnvironmentArnParameters(scope constructs.Construct, id *string, computeEnvironmentArn *string) error {
	return nil
}

func validateComputeEnvironment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateComputeEnvironment_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewComputeEnvironmentParameters(scope constructs.Construct, id *string, props *ComputeEnvironmentProps) error {
	return nil
}

