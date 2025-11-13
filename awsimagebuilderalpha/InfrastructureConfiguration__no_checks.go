//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InfrastructureConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_InfrastructureConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_InfrastructureConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureConfiguration) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_InfrastructureConfiguration) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateInfrastructureConfiguration_FromInfrastructureConfigurationArnParameters(scope constructs.Construct, id *string, infrastructureConfigurationArn *string) error {
	return nil
}

func validateInfrastructureConfiguration_FromInfrastructureConfigurationNameParameters(scope constructs.Construct, id *string, infrastructureConfigurationName *string) error {
	return nil
}

func validateInfrastructureConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInfrastructureConfiguration_IsInfrastructureConfigurationParameters(x interface{}) error {
	return nil
}

func validateInfrastructureConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateInfrastructureConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewInfrastructureConfigurationParameters(scope constructs.Construct, id *string, props *InfrastructureConfigurationProps) error {
	return nil
}

