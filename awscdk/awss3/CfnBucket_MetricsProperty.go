package awss3


// A container specifying replication metrics-related settings enabling replication metrics and events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsProperty := &MetricsProperty{
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	EventThreshold: &ReplicationTimeValueProperty{
//   		Minutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metrics.html
//
type CfnBucket_MetricsProperty struct {
	// Specifies whether the replication metrics are enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metrics.html#cfn-s3-bucket-metrics-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// A container specifying the time threshold for emitting the `s3:Replication:OperationMissedThreshold` event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metrics.html#cfn-s3-bucket-metrics-eventthreshold
	//
	EventThreshold interface{} `field:"optional" json:"eventThreshold" yaml:"eventThreshold"`
}

