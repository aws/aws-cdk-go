//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func validateFilterOrPolicy_FilterParameters(filter SubscriptionFilter) error {
	return nil
}

func validateFilterOrPolicy_PolicyParameters(policy *map[string]FilterOrPolicy) error {
	return nil
}

