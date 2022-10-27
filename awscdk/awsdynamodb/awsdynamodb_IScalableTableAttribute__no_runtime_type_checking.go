//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IScalableTableAttribute) validateScaleOnScheduleParameters(id *string, actions *awsapplicationautoscaling.ScalingSchedule) error {
	return nil
}

func (i *jsiiProxy_IScalableTableAttribute) validateScaleOnUtilizationParameters(props *UtilizationScalingProps) error {
	return nil
}

