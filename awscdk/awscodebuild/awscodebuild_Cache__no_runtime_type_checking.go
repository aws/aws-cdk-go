//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func validateCache_BucketParameters(bucket awss3.IBucket, options *BucketCacheOptions) error {
	return nil
}

