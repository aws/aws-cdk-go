package awsimagebuilder


// Amazon S3 logging configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LogsProperty := &s3LogsProperty{
//   	s3BucketName: jsii.String("s3BucketName"),
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
type CfnInfrastructureConfiguration_S3LogsProperty struct {
	// The S3 bucket in which to store the logs.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// The Amazon S3 path to the bucket where the logs are stored.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

