package awsecs


// The log configuration for the results of the execute command actions.
//
// The logs can be sent to CloudWatch Logs or an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executeCommandLogConfigurationProperty := &executeCommandLogConfigurationProperty{
//   	cloudWatchEncryptionEnabled: jsii.Boolean(false),
//   	cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   	s3BucketName: jsii.String("s3BucketName"),
//   	s3EncryptionEnabled: jsii.Boolean(false),
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
type CfnCluster_ExecuteCommandLogConfigurationProperty struct {
	// Determines whether to use encryption on the CloudWatch logs.
	//
	// If not specified, encryption will be disabled.
	CloudWatchEncryptionEnabled interface{} `field:"optional" json:"cloudWatchEncryptionEnabled" yaml:"cloudWatchEncryptionEnabled"`
	// The name of the CloudWatch log group to send logs to.
	//
	// > The CloudWatch log group must already be created.
	CloudWatchLogGroupName *string `field:"optional" json:"cloudWatchLogGroupName" yaml:"cloudWatchLogGroupName"`
	// The name of the S3 bucket to send logs to.
	//
	// > The S3 bucket must already be created.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Determines whether to use encryption on the S3 logs.
	//
	// If not specified, encryption is not used.
	S3EncryptionEnabled interface{} `field:"optional" json:"s3EncryptionEnabled" yaml:"s3EncryptionEnabled"`
	// An optional folder in the S3 bucket to place logs in.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

