package awss3tables


// The key format of the log objects.
//
// Only PartitionedPrefix is supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetObjectKeyFormatProperty := &TargetObjectKeyFormatProperty{
//   	PartitionedPrefix: &PartitionedPrefixProperty{
//   		PartitionDateSource: jsii.String("partitionDateSource"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-targetobjectkeyformat.html
//
type CfnTableBucket_TargetObjectKeyFormatProperty struct {
	// Amazon S3 keys for log objects are partitioned in the following format: [TargetPrefix][SourceAccountId]/[SourceRegion]/[SourceTableBucket]/[YYYY]/[MM]/[DD]/[YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString].
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-targetobjectkeyformat.html#cfn-s3tables-tablebucket-targetobjectkeyformat-partitionedprefix
	//
	PartitionedPrefix interface{} `field:"required" json:"partitionedPrefix" yaml:"partitionedPrefix"`
}

