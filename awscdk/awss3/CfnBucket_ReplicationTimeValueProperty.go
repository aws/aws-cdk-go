package awss3


// A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics `EventThreshold` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationTimeValueProperty := &ReplicationTimeValueProperty{
//   	Minutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationtimevalue.html
//
type CfnBucket_ReplicationTimeValueProperty struct {
	// Contains an integer specifying time in minutes.
	//
	// Valid value: 15.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationtimevalue.html#cfn-s3-bucket-replicationtimevalue-minutes
	//
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

