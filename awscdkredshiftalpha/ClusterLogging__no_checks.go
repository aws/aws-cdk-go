//go:build no_runtime_type_checking

package awscdkredshiftalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateClusterLogging_CloudWatchParameters(options *CloudWatchLoggingOptions) error {
	return nil
}

func validateClusterLogging_S3Parameters(options *S3LoggingOptions) error {
	return nil
}

