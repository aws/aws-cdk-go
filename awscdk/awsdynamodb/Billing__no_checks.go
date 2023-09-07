//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func validateBilling_ProvisionedParameters(props *ThroughputProps) error {
	return nil
}

