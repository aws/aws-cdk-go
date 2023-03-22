//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateMultipartBody_FromRawBodyParameters(opts *MultipartBodyOptions) error {
	return nil
}

func validateMultipartBody_FromUserDataParameters(userData UserData) error {
	return nil
}

