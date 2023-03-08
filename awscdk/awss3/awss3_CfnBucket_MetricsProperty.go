package awss3


// A container specifying replication metrics-related settings enabling replication metrics and events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsProperty := &metricsProperty{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	eventThreshold: &replicationTimeValueProperty{
//   		minutes: jsii.Number(123),
//   	},
//   }
//
type CfnBucket_MetricsProperty struct {
	// Specifies whether the replication metrics are enabled.
	Status *string `field:"required" json:"status" yaml:"status"`
	// A container specifying the time threshold for emitting the `s3:Replication:OperationMissedThreshold` event.
	EventThreshold interface{} `field:"optional" json:"eventThreshold" yaml:"eventThreshold"`
}

