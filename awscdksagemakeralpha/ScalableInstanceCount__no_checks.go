//go:build no_runtime_type_checking

package awscdksagemakeralpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScalableInstanceCount) validateDoScaleOnMetricParameters(id *string, props *awsapplicationautoscaling.BasicStepScalingPolicyProps) error {
	return nil
}

func (s *jsiiProxy_ScalableInstanceCount) validateDoScaleOnScheduleParameters(id *string, props *awsapplicationautoscaling.ScalingSchedule) error {
	return nil
}

func (s *jsiiProxy_ScalableInstanceCount) validateDoScaleToTrackMetricParameters(id *string, props *awsapplicationautoscaling.BasicTargetTrackingScalingPolicyProps) error {
	return nil
}

func (s *jsiiProxy_ScalableInstanceCount) validateScaleOnInvocationsParameters(id *string, props *InvocationsScalingProps) error {
	return nil
}

func validateScalableInstanceCount_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewScalableInstanceCountParameters(scope constructs.Construct, id *string, props *ScalableInstanceCountProps) error {
	return nil
}

