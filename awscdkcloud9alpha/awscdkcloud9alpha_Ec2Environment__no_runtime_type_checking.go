//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Cloud9
package awscdkcloud9alpha

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

func validateEc2Environment_FromEc2EnvironmentNameParameters(scope constructs.Construct, id *string, ec2EnvironmentName *string) error {
	return nil
}

func validateEc2Environment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEc2Environment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEc2Environment_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEc2EnvironmentParameters(scope constructs.Construct, id *string, props *Ec2EnvironmentProps) error {
	return nil
}

