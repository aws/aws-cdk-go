package awscloudwatch


// A reference to a MetricStream resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricStreamReference := &MetricStreamReference{
//   	MetricStreamArn: jsii.String("metricStreamArn"),
//   	MetricStreamName: jsii.String("metricStreamName"),
//   }
//
type MetricStreamReference struct {
	// The ARN of the MetricStream resource.
	MetricStreamArn *string `field:"required" json:"metricStreamArn" yaml:"metricStreamArn"`
	// The Name of the MetricStream resource.
	MetricStreamName *string `field:"required" json:"metricStreamName" yaml:"metricStreamName"`
}

