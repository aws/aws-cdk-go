package awsec2


// Describes the options for Verified Access logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   verifiedAccessLogsProperty := &VerifiedAccessLogsProperty{
//   	CloudWatchLogs: &CloudWatchLogsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	IncludeTrustContext: jsii.Boolean(false),
//   	KinesisDataFirehose: &KinesisDataFirehoseProperty{
//   		DeliveryStream: jsii.String("deliveryStream"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   	LogVersion: jsii.String("logVersion"),
//   	S3: &S3Property{
//   		BucketName: jsii.String("bucketName"),
//   		BucketOwner: jsii.String("bucketOwner"),
//   		Enabled: jsii.Boolean(false),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesslogs.html
//
type CfnVerifiedAccessInstance_VerifiedAccessLogsProperty struct {
	// CloudWatch Logs logging destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesslogs.html#cfn-ec2-verifiedaccessinstance-verifiedaccesslogs-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Indicates whether to include trust data sent by trust providers in the logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesslogs.html#cfn-ec2-verifiedaccessinstance-verifiedaccesslogs-includetrustcontext
	//
	IncludeTrustContext interface{} `field:"optional" json:"includeTrustContext" yaml:"includeTrustContext"`
	// Kinesis logging destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesslogs.html#cfn-ec2-verifiedaccessinstance-verifiedaccesslogs-kinesisdatafirehose
	//
	KinesisDataFirehose interface{} `field:"optional" json:"kinesisDataFirehose" yaml:"kinesisDataFirehose"`
	// The logging version.
	//
	// Valid values: `ocsf-0.1` | `ocsf-1.0.0-rc.2`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesslogs.html#cfn-ec2-verifiedaccessinstance-verifiedaccesslogs-logversion
	//
	LogVersion *string `field:"optional" json:"logVersion" yaml:"logVersion"`
	// Amazon S3 logging options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesslogs.html#cfn-ec2-verifiedaccessinstance-verifiedaccesslogs-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

