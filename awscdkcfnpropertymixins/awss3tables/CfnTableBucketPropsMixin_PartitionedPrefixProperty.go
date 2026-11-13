package awss3tables


// Amazon S3 keys for log objects are partitioned in the following format: [TargetPrefix][SourceAccountId]/[SourceRegion]/[SourceTableBucket]/[YYYY]/[MM]/[DD]/[YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString].
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   partitionedPrefixProperty := &PartitionedPrefixProperty{
//   	PartitionDateSource: jsii.String("partitionDateSource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-partitionedprefix.html
//
type CfnTableBucketPropsMixin_PartitionedPrefixProperty struct {
	// The date to use for the log object key partitions.
	//
	// Valid values are EventTime, the time that the requests in the log were made, and DeliveryTime, the time that the log was delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-partitionedprefix.html#cfn-s3tables-tablebucket-partitionedprefix-partitiondatesource
	//
	PartitionDateSource *string `field:"optional" json:"partitionDateSource" yaml:"partitionDateSource"`
}

