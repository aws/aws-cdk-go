package awss3tables


// Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a table bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationProperty := &LoggingConfigurationProperty{
//   	TargetBucket: jsii.String("targetBucket"),
//   	TargetObjectKeyFormat: &TargetObjectKeyFormatProperty{
//   		PartitionedPrefix: &PartitionedPrefixProperty{
//   			PartitionDateSource: jsii.String("partitionDateSource"),
//   		},
//   	},
//   	TargetPrefix: jsii.String("targetPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-loggingconfiguration.html
//
type CfnTableBucket_LoggingConfigurationProperty struct {
	// The name of the general purpose bucket that logs are delivered to.
	//
	// The bucket must be in the same AWS Region and owned by the same AWS account as the table bucket. The bucket can't use SSE-KMS or DSSE-KMS default encryption, can't have S3 Object Lock enabled, and can't have Requester Pays enabled. You can deliver logs from multiple table buckets to the same target bucket. In this case, choose a different TargetPrefix for each table bucket so that you can tell their log objects apart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-loggingconfiguration.html#cfn-s3tables-tablebucket-loggingconfiguration-targetbucket
	//
	TargetBucket *string `field:"required" json:"targetBucket" yaml:"targetBucket"`
	// The key format of the log objects.
	//
	// Only PartitionedPrefix is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-loggingconfiguration.html#cfn-s3tables-tablebucket-loggingconfiguration-targetobjectkeyformat
	//
	TargetObjectKeyFormat interface{} `field:"required" json:"targetObjectKeyFormat" yaml:"targetObjectKeyFormat"`
	// The prefix for the keys of the log objects that are delivered to the target bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-loggingconfiguration.html#cfn-s3tables-tablebucket-loggingconfiguration-targetprefix
	//
	TargetPrefix *string `field:"required" json:"targetPrefix" yaml:"targetPrefix"`
}

