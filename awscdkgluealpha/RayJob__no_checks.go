//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RayJob) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateBuildJobArnParameters(scope constructs.Construct, jobName *string) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateCodeS3ObjectUrlParameters(code Code) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateMetricParameters(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateMetricFailureParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateMetricSuccessParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateMetricTimeoutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateOnFailureParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateOnStateChangeParameters(id *string, jobState JobState, options *awsevents.OnEventOptions) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateOnSuccessParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateOnTimeoutParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (r *jsiiProxy_RayJob) validateSetupContinuousLoggingParameters(role awsiam.IRole, props *ContinuousLoggingProps) error {
	return nil
}

func validateRayJob_FromJobAttributesParameters(scope constructs.Construct, id *string, attrs *JobAttributes) error {
	return nil
}

func validateRayJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRayJob_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRayJob_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRayJobParameters(scope constructs.Construct, id *string, props *RayJobProps) error {
	return nil
}

