//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK Destinations Constructs for AWS Kinesis Firehose
package awscdkkinesisfirehosedestinationsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Bucket) validateBindParameters(scope constructs.Construct, _options *awscdkkinesisfirehosealpha.DestinationBindOptions) error {
	return nil
}

func validateNewS3BucketParameters(bucket awss3.IBucket, props *S3BucketProps) error {
	return nil
}

