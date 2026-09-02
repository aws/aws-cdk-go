//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateOutputDestination_ToBucketParameters(bucket awss3.IBucket) error {
	return nil
}

func validateOutputDestination_UrlParameters(url *string, options *OutputDestinationOptions) error {
	return nil
}

