//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScalableTaskCount) validateDoScaleOnMetricParameters(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) error {
	return nil
}

func (s *jsiiProxy_ScalableTaskCount) validateDoScaleOnScheduleParameters(id *string, props *awsapplicationautoscaling.ScalingSchedule) error {
	return nil
}

func (s *jsiiProxy_ScalableTaskCount) validateDoScaleToTrackMetricParameters(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps) error {
	return nil
}

func (s *jsiiProxy_ScalableTaskCount) validateScaleOnCpuUtilizationParameters(id *string, props *CpuUtilizationScalingProps) error {
	return nil
}

func (s *jsiiProxy_ScalableTaskCount) validateScaleOnMemoryUtilizationParameters(id *string, props *MemoryUtilizationScalingProps) error {
	return nil
}

func (s *jsiiProxy_ScalableTaskCount) validateScaleOnMetricParameters(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) error {
	return nil
}

func (s *jsiiProxy_ScalableTaskCount) validateScaleOnRequestCountParameters(id *string, props *RequestCountScalingProps) error {
	return nil
}

func (s *jsiiProxy_ScalableTaskCount) validateScaleOnScheduleParameters(id *string, props *awsapplicationautoscaling.ScalingSchedule) error {
	return nil
}

func (s *jsiiProxy_ScalableTaskCount) validateScaleToTrackCustomMetricParameters(id *string, props *TrackCustomMetricProps) error {
	return nil
}

func validateScalableTaskCount_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewScalableTaskCountParameters(scope constructs.Construct, id *string, props *ScalableTaskCountProps) error {
	return nil
}

