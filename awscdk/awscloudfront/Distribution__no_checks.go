//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Distribution) validateAddBehaviorParameters(pathPattern *string, origin IOrigin, behaviorOptions *AddBehaviorOptions) error {
	return nil
}

func (d *jsiiProxy_Distribution) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_Distribution) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_Distribution) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_Distribution) validateGrantParameters(identity awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_Distribution) validateGrantCreateInvalidationParameters(identity awsiam.IGrantable) error {
	return nil
}

func validateDistribution_FromDistributionAttributesParameters(scope constructs.Construct, id *string, attrs *DistributionAttributes) error {
	return nil
}

func validateDistribution_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDistribution_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDistribution_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDistributionParameters(scope constructs.Construct, id *string, props *DistributionProps) error {
	return nil
}

