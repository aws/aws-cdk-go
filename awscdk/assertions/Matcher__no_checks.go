//go:build no_runtime_type_checking

package assertions

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Matcher) validateTestParameters(actual interface{}) error {
	return nil
}

func validateMatcher_IsMatcherParameters(x interface{}) error {
	return nil
}

