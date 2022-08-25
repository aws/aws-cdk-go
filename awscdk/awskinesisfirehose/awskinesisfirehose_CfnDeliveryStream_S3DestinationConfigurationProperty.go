package awskinesisfirehose


// The `S3DestinationConfiguration` property type specifies an Amazon Simple Storage Service (Amazon S3) destination to which Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationConfigurationProperty := &s3DestinationConfigurationProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	bufferingHints: &bufferingHintsProperty{
//   		intervalInSeconds: jsii.Number(123),
//   		sizeInMBs: jsii.Number(123),
//   	},
//   	cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		logGroupName: jsii.String("logGroupName"),
//   		logStreamName: jsii.String("logStreamName"),
//   	},
//   	compressionFormat: jsii.String("compressionFormat"),
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   			awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   		},
//   		noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   	},
//   	errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnDeliveryStream_S3DestinationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket to send data to.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The ARN of an AWS Identity and Access Management (IAM) role that grants Kinesis Data Firehose access to your Amazon S3 bucket and AWS KMS (if you enable data encryption).
	//
	// For more information, see [Grant Kinesis Data Firehose Access to an Amazon S3 Destination](https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html#using-iam-s3) in the *Amazon Kinesis Data Firehose Developer Guide* .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Configures how Kinesis Data Firehose buffers incoming data while delivering it to the Amazon S3 bucket.
	BufferingHints interface{} `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// The CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// For valid values, see the `CompressionFormat` content for the [S3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_S3DestinationConfiguration.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	CompressionFormat *string `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
	// Configures Amazon Simple Storage Service (Amazon S3) server-side encryption.
	//
	// Kinesis Data Firehose uses AWS Key Management Service ( AWS KMS) to encrypt the data that it delivers to your Amazon S3 bucket.
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html) .
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// A prefix that Kinesis Data Firehose adds to the files that it delivers to the Amazon S3 bucket.
	//
	// The prefix helps you identify the files that Kinesis Data Firehose delivered.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

