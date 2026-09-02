//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateLut_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateLut_UrlParameters(url *string) error {
	return nil
}

