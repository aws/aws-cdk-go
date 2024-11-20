package awslambda


// Metrics config for Event Source Mapping Metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsConfigProperty := &MetricsConfigProperty{
//   	Metrics: []*string{
//   		jsii.String("metrics"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-metricsconfig.html
//
type CfnEventSourceMapping_MetricsConfigProperty struct {
	// Metric groups to enable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-metricsconfig.html#cfn-lambda-eventsourcemapping-metricsconfig-metrics
	//
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
}

