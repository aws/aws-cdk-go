//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func validateFilter_FilterParameters(filter SubscriptionFilter) error {
	return nil
}

func validateFilter_PolicyParameters(policy *map[string]FilterOrPolicy) error {
	return nil
}

func validateNewFilterParameters(filterDoc SubscriptionFilter) error {
	return nil
}

