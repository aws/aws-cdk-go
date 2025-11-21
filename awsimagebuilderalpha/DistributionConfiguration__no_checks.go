//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DistributionConfiguration) validateAddAmiDistributionsParameters(amiDistributions *[]*AmiDistribution) error {
	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateAddContainerDistributionsParameters(containerDistributions *[]*ContainerDistribution) error {
	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateDistributionConfiguration_FromDistributionConfigurationArnParameters(scope constructs.Construct, id *string, distributionConfigurationArn *string) error {
	return nil
}

func validateDistributionConfiguration_FromDistributionConfigurationNameParameters(scope constructs.Construct, id *string, distributionConfigurationName *string) error {
	return nil
}

func validateDistributionConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDistributionConfiguration_IsDistributionConfigurationParameters(x interface{}) error {
	return nil
}

func validateDistributionConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDistributionConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDistributionConfigurationParameters(scope constructs.Construct, id *string, props *DistributionConfigurationProps) error {
	return nil
}

