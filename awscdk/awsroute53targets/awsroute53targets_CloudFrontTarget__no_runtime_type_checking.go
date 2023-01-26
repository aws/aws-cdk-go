//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFrontTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateCloudFrontTarget_GetHostedZoneIdParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewCloudFrontTargetParameters(distribution awscloudfront.IDistribution) error {
	return nil
}

