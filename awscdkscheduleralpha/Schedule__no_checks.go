//go:build no_runtime_type_checking

package awscdkscheduleralpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Schedule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Schedule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Schedule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSchedule_FromScheduleArnParameters(scope constructs.Construct, id *string, scheduleArn *string) error {
	return nil
}

func validateSchedule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSchedule_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSchedule_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSchedule_MetricAllParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateSchedule_MetricAllAttemptsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateSchedule_MetricAllDroppedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateSchedule_MetricAllErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateSchedule_MetricAllFailedToBeSentToDLQParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateSchedule_MetricAllSentToDLQParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateSchedule_MetricAllSentToDLQTruncatedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateSchedule_MetricAllTargetThrottledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateSchedule_MetricAllThrottledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNewScheduleParameters(scope constructs.Construct, id *string, props *ScheduleProps) error {
	return nil
}

