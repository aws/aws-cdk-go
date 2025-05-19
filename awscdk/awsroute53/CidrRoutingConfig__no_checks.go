//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func validateCidrRoutingConfig_CreateParameters(props *CidrRoutingConfigProps) error {
	return nil
}

func validateCidrRoutingConfig_WithDefaultLocationNameParameters(collectionId *string) error {
	return nil
}

