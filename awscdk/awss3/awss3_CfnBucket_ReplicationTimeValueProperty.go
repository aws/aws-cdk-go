package awss3


// A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics `EventThreshold` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationTimeValueProperty := &replicationTimeValueProperty{
//   	minutes: jsii.Number(123),
//   }
//
type CfnBucket_ReplicationTimeValueProperty struct {
	// Contains an integer specifying time in minutes.
	//
	// Valid value: 15.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

