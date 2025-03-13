//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateBuildJobArnParameters(scope constructs.Construct, jobName *string) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateMetricParameters(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateMetricFailureParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateMetricSuccessParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateMetricTimeoutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateOnFailureParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateOnStateChangeParameters(id *string, jobState JobState, options *awsevents.OnEventOptions) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateOnSuccessParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (j *jsiiProxy_JobBase) validateOnTimeoutParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateJobBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateJobBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateJobBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewJobBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

