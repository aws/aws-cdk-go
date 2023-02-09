//go:build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Job) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateMetricParameters(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (j *jsiiProxy_Job) validateMetricFailureParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (j *jsiiProxy_Job) validateMetricSuccessParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (j *jsiiProxy_Job) validateMetricTimeoutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (j *jsiiProxy_Job) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (j *jsiiProxy_Job) validateOnFailureParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (j *jsiiProxy_Job) validateOnStateChangeParameters(id *string, jobState JobState, options *awsevents.OnEventOptions) error {
	return nil
}

func (j *jsiiProxy_Job) validateOnSuccessParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (j *jsiiProxy_Job) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (j *jsiiProxy_Job) validateOnTimeoutParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (j *jsiiProxy_Job) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateJob_FromJobAttributesParameters(scope constructs.Construct, id *string, attrs *JobAttributes) error {
	return nil
}

func validateJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateJob_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewJobParameters(scope constructs.Construct, id *string, props *JobProps) error {
	return nil
}

