//go:build no_runtime_type_checking

package awsscheduler

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IScheduleGroup) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateGrantDeleteSchedulesParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateGrantReadSchedulesParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateGrantWriteSchedulesParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricAttemptsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricDroppedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricFailedToBeSentToDLQParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricSentToDLQParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricSentToDLQTruncatedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricTargetErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricTargetThrottledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IScheduleGroup) validateMetricThrottledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

