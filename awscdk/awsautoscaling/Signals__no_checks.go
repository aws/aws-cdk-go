//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Signals) validateDoRenderParameters(options *SignalsOptions) error {
	return nil
}

func (s *jsiiProxy_Signals) validateRenderCreationPolicyParameters(renderOptions *RenderSignalsOptions) error {
	return nil
}

func validateSignals_WaitForAllParameters(options *SignalsOptions) error {
	return nil
}

func validateSignals_WaitForCountParameters(count *float64, options *SignalsOptions) error {
	return nil
}

func validateSignals_WaitForMinCapacityParameters(options *SignalsOptions) error {
	return nil
}

