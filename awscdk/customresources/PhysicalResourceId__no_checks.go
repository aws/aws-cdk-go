//go:build no_runtime_type_checking

package customresources

// Building without runtime type checking enabled, so all the below just return nil

func validatePhysicalResourceId_FromResponseParameters(responsePath *string) error {
	return nil
}

func validatePhysicalResourceId_OfParameters(id *string) error {
	return nil
}

