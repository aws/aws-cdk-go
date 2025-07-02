//go:build no_runtime_type_checking

package awsapplicationautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BaseScalableAttribute) validateDoScaleOnMetricParameters(id *string, props *BasicStepScalingPolicyProps) error {
	return nil
}

func (b *jsiiProxy_BaseScalableAttribute) validateDoScaleOnScheduleParameters(id *string, props *ScalingSchedule) error {
	return nil
}

func (b *jsiiProxy_BaseScalableAttribute) validateDoScaleToTrackMetricParameters(id *string, props *BasicTargetTrackingScalingPolicyProps) error {
	return nil
}

func validateBaseScalableAttribute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewBaseScalableAttributeParameters(scope constructs.Construct, id *string, props *BaseScalableAttributeProps) error {
	return nil
}

