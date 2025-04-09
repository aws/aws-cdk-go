//go:build no_runtime_type_checking

package awsautoscalingcommon

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IRandomGenerator) validateNextIntParameters(min *float64, max *float64) error {
	return nil
}

