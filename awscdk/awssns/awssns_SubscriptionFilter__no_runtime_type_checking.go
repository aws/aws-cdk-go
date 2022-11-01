//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func validateSubscriptionFilter_NumericFilterParameters(numericConditions *NumericConditions) error {
	return nil
}

func validateSubscriptionFilter_StringFilterParameters(stringConditions *StringConditions) error {
	return nil
}

