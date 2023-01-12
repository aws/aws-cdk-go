//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Activity) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Activity) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Activity) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_Activity) validateGrantParameters(identity awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_Activity) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Activity) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Activity) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Activity) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Activity) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Activity) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Activity) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Activity) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Activity) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Activity) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateActivity_FromActivityArnParameters(scope constructs.Construct, id *string, activityArn *string) error {
	return nil
}

func validateActivity_FromActivityNameParameters(scope constructs.Construct, id *string, activityName *string) error {
	return nil
}

func validateActivity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateActivity_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateActivity_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewActivityParameters(scope constructs.Construct, id *string, props *ActivityProps) error {
	return nil
}

