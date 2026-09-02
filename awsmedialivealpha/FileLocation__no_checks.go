//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateFileLocation_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateFileLocation_UrlParameters(url *string, options *FileLocationOptions) error {
	return nil
}

