//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func validateCapacityProviderOptions_CustomParameters(capacityProviderStrategy *[]*awsecs.CapacityProviderStrategy) error {
	return nil
}

