//go:build no_runtime_type_checking

package awsapplicationautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepScalingPolicy) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_StepScalingPolicy) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateStepScalingPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewStepScalingPolicyParameters(scope constructs.Construct, id *string, props *StepScalingPolicyProps) error {
	return nil
}

