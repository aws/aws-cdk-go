package awss3


// Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a bucket.
//
// For examples and more information, see [PUT Bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html) in the *Amazon S3 API Reference* .
//
// > To successfully complete the `AWS::S3::Bucket LoggingConfiguration` request, you must have `s3:PutObject` and `s3:PutObjectAcl` in your IAM permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var simplePrefix interface{}
//
//   loggingConfigurationProperty := &LoggingConfigurationProperty{
//   	DestinationBucketName: jsii.String("destinationBucketName"),
//   	LogFilePrefix: jsii.String("logFilePrefix"),
//   	TargetObjectKeyFormat: &TargetObjectKeyFormatProperty{
//   		PartitionedPrefix: &PartitionedPrefixProperty{
//   			PartitionDateSource: jsii.String("partitionDateSource"),
//   		},
//   		SimplePrefix: simplePrefix,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfiguration.html
//
type CfnBucket_LoggingConfigurationProperty struct {
	// The name of the bucket where Amazon S3 should store server access log files.
	//
	// You can store log files in any bucket that you own. By default, logs are stored in the bucket where the `LoggingConfiguration` property is defined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfiguration.html#cfn-s3-bucket-loggingconfiguration-destinationbucketname
	//
	DestinationBucketName *string `field:"optional" json:"destinationBucketName" yaml:"destinationBucketName"`
	// A prefix for all log object keys.
	//
	// If you store log files from multiple Amazon S3 buckets in a single bucket, you can use a prefix to distinguish which log files came from which bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfiguration.html#cfn-s3-bucket-loggingconfiguration-logfileprefix
	//
	LogFilePrefix *string `field:"optional" json:"logFilePrefix" yaml:"logFilePrefix"`
	// Amazon S3 key format for log objects.
	//
	// Only one format, either PartitionedPrefix or SimplePrefix, is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfiguration.html#cfn-s3-bucket-loggingconfiguration-targetobjectkeyformat
	//
	TargetObjectKeyFormat interface{} `field:"optional" json:"targetObjectKeyFormat" yaml:"targetObjectKeyFormat"`
}

