//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IStateMachine) validateGrantParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateGrantExecutionParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateGrantReadParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateGrantStartExecutionParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateGrantStartSyncExecutionParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateGrantTaskResponseParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateMetricAbortedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateMetricThrottledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStateMachine) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

