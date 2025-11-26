package previewawslambdamixins


// The metrics configuration for your event source.
//
// Use this configuration object to define which metrics you want your event source mapping to produce.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricsConfigProperty := &MetricsConfigProperty{
//   	Metrics: []*string{
//   		jsii.String("metrics"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-metricsconfig.html
//
type CfnEventSourceMappingPropsMixin_MetricsConfigProperty struct {
	// The metrics you want your event source mapping to produce.
	//
	// Include `EventCount` to receive event source mapping metrics related to the number of events processed by your event source mapping. For more information about these metrics, see [Event source mapping metrics](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics-types.html#event-source-mapping-metrics) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-metricsconfig.html#cfn-lambda-eventsourcemapping-metricsconfig-metrics
	//
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
}

