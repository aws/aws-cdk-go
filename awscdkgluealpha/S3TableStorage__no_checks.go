//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateS3TableStorage_FromBucketParameters(bucket awss3.IBucket) error {
	return nil
}

