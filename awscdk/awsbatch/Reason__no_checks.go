//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func validateReason_CustomParameters(customReasonProps *CustomReason) error {
	return nil
}

