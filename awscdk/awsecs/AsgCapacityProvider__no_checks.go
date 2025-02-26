//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func validateAsgCapacityProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAsgCapacityProviderParameters(scope constructs.Construct, id *string, props *AsgCapacityProviderProps) error {
	return nil
}

