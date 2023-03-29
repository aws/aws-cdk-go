//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFrontWebDistribution) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CloudFrontWebDistribution) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CloudFrontWebDistribution) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CloudFrontWebDistribution) validateGrantParameters(identity awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_CloudFrontWebDistribution) validateGrantCreateInvalidationParameters(identity awsiam.IGrantable) error {
	return nil
}

func validateCloudFrontWebDistribution_FromDistributionAttributesParameters(scope constructs.Construct, id *string, attrs *CloudFrontWebDistributionAttributes) error {
	return nil
}

func validateCloudFrontWebDistribution_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCloudFrontWebDistribution_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCloudFrontWebDistribution_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCloudFrontWebDistributionParameters(scope constructs.Construct, id *string, props *CloudFrontWebDistributionProps) error {
	return nil
}

