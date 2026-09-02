//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateS3OutputDestination_ToBucketParameters(bucket awss3.IBucket) error {
	return nil
}

func validateS3OutputDestination_UrlParameters(url *string) error {
	return nil
}

