package awss3


// A container specifying S3 Replication Time Control (S3 RTC) related information, including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated.
//
// Must be specified together with a `Metrics` block.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationTimeProperty := &ReplicationTimeProperty{
//   	Status: jsii.String("status"),
//   	Time: &ReplicationTimeValueProperty{
//   		Minutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationtime.html
//
type CfnBucket_ReplicationTimeProperty struct {
	// Specifies whether the replication time is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationtime.html#cfn-s3-bucket-replicationtime-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// A container specifying the time by which replication should be complete for all objects and operations on objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationtime.html#cfn-s3-bucket-replicationtime-time
	//
	Time interface{} `field:"required" json:"time" yaml:"time"`
}

