package awss3


// Amazon S3 keys for log objects are partitioned in the following format:.
//
// `[DestinationPrefix][SourceAccountId]/[SourceRegion]/[SourceBucket]/[YYYY]/[MM]/[DD]/[YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString]`
//
// PartitionedPrefix defaults to EventTime delivery when server access logs are delivered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partitionedPrefixProperty := &PartitionedPrefixProperty{
//   	PartitionDateSource: jsii.String("partitionDateSource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-partitionedprefix.html
//
type CfnBucket_PartitionedPrefixProperty struct {
	// Specifies the partition date source for the partitioned prefix. `PartitionDateSource` can be `EventTime` or `DeliveryTime` .
	//
	// For `DeliveryTime` , the time in the log file names corresponds to the delivery time for the log files.
	//
	// For `EventTime` , The logs delivered are for a specific day only. The year, month, and day correspond to the day on which the event occurred, and the hour, minutes and seconds are set to 00 in the key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-partitionedprefix.html#cfn-s3-bucket-partitionedprefix-partitiondatesource
	//
	PartitionDateSource *string `field:"optional" json:"partitionDateSource" yaml:"partitionDateSource"`
}

