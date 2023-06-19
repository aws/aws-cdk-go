//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateInitService_DisableParameters(serviceName *string) error {
	return nil
}

func validateInitService_EnableParameters(serviceName *string, options *InitServiceOptions) error {
	return nil
}

