//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscloud9

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2Environment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_Ec2Environment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_Ec2Environment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_Ec2Environment) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (e *jsiiProxy_Ec2Environment) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateEc2Environment_FromEc2EnvironmentNameParameters(scope constructs.Construct, id *string, ec2EnvironmentName *string) error {
	return nil
}

func validateEc2Environment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEc2Environment_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewEc2EnvironmentParameters(scope constructs.Construct, id *string, props *Ec2EnvironmentProps) error {
	return nil
}

