//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAutoScalingGroup) validateAddLifecycleHookParameters(id *string, props *BasicLifecycleHookProps) error {
	return nil
}

func (i *jsiiProxy_IAutoScalingGroup) validateAddWarmPoolParameters(options *WarmPoolOptions) error {
	return nil
}

func (i *jsiiProxy_IAutoScalingGroup) validateScaleOnCpuUtilizationParameters(id *string, props *CpuUtilizationScalingProps) error {
	return nil
}

func (i *jsiiProxy_IAutoScalingGroup) validateScaleOnIncomingBytesParameters(id *string, props *NetworkUtilizationScalingProps) error {
	return nil
}

func (i *jsiiProxy_IAutoScalingGroup) validateScaleOnMetricParameters(id *string, props *BasicStepScalingPolicyProps) error {
	return nil
}

func (i *jsiiProxy_IAutoScalingGroup) validateScaleOnOutgoingBytesParameters(id *string, props *NetworkUtilizationScalingProps) error {
	return nil
}

func (i *jsiiProxy_IAutoScalingGroup) validateScaleOnScheduleParameters(id *string, props *BasicScheduledActionProps) error {
	return nil
}

func (i *jsiiProxy_IAutoScalingGroup) validateScaleToTrackMetricParameters(id *string, props *MetricTargetTrackingProps) error {
	return nil
}

func (i *jsiiProxy_IAutoScalingGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

