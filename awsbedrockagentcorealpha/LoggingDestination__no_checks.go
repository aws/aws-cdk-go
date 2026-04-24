//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateLoggingDestination_CloudWatchLogsParameters(logGroup awslogs.ILogGroup) error {
	return nil
}

func validateLoggingDestination_FirehoseParameters(stream awskinesisfirehose.IDeliveryStream) error {
	return nil
}

func validateLoggingDestination_S3Parameters(bucket awss3.IBucket) error {
	return nil
}

