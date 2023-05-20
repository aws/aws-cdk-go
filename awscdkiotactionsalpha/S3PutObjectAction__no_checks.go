//go:build no_runtime_type_checking

package awscdkiotactionsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateNewS3PutObjectActionParameters(bucket awss3.IBucket, props *S3PutObjectActionProps) error {
	return nil
}

