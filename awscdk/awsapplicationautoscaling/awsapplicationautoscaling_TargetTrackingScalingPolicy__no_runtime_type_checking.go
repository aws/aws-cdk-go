//go:build no_runtime_type_checking

package awsapplicationautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TargetTrackingScalingPolicy) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (t *jsiiProxy_TargetTrackingScalingPolicy) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateTargetTrackingScalingPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTargetTrackingScalingPolicyParameters(scope constructs.Construct, id *string, props *TargetTrackingScalingPolicyProps) error {
	return nil
}

