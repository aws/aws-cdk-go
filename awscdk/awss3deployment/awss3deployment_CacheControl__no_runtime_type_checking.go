//go:build no_runtime_type_checking

package awss3deployment

// Building without runtime type checking enabled, so all the below just return nil

func validateCacheControl_FromStringParameters(s *string) error {
	return nil
}

func validateCacheControl_MaxAgeParameters(t awscdk.Duration) error {
	return nil
}

func validateCacheControl_SMaxAgeParameters(t awscdk.Duration) error {
	return nil
}

