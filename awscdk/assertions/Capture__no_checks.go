//go:build no_runtime_type_checking

package assertions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Capture) validateTestParameters(actual interface{}) error {
	return nil
}

func validateCapture_IsMatcherParameters(x interface{}) error {
	return nil
}

