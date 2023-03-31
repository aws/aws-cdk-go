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
//   replicationTimeProperty := &replicationTimeProperty{
//   	status: jsii.String("status"),
//   	time: &replicationTimeValueProperty{
//   		minutes: jsii.Number(123),
//   	},
//   }
//
type CfnBucket_ReplicationTimeProperty struct {
	// Specifies whether the replication time is enabled.
	Status *string `field:"required" json:"status" yaml:"status"`
	// A container specifying the time by which replication should be complete for all objects and operations on objects.
	Time interface{} `field:"required" json:"time" yaml:"time"`
}

