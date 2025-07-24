package awss3


// Amazon S3 key format for log objects.
//
// Only one format, PartitionedPrefix or SimplePrefix, is allowed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var simplePrefix interface{}
//
//   targetObjectKeyFormatProperty := &TargetObjectKeyFormatProperty{
//   	PartitionedPrefix: &PartitionedPrefixProperty{
//   		PartitionDateSource: jsii.String("partitionDateSource"),
//   	},
//   	SimplePrefix: simplePrefix,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-targetobjectkeyformat.html
//
type CfnBucket_TargetObjectKeyFormatProperty struct {
	// Partitioned S3 key for log objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-targetobjectkeyformat.html#cfn-s3-bucket-targetobjectkeyformat-partitionedprefix
	//
	PartitionedPrefix interface{} `field:"optional" json:"partitionedPrefix" yaml:"partitionedPrefix"`
	// To use the simple format for S3 keys for log objects.
	//
	// To specify SimplePrefix format, set SimplePrefix to {}.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-targetobjectkeyformat.html#cfn-s3-bucket-targetobjectkeyformat-simpleprefix
	//
	SimplePrefix interface{} `field:"optional" json:"simplePrefix" yaml:"simplePrefix"`
}

