//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func validateStepScalingPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewStepScalingPolicyParameters(scope constructs.Construct, id *string, props *StepScalingPolicyProps) error {
	return nil
}

