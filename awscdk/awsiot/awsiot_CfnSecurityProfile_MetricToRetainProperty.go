package awsiot


// The metric you want to retain.
//
// Dimensions are optional.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricToRetainProperty := &metricToRetainProperty{
//   	metric: jsii.String("metric"),
//
//   	// the properties below are optional
//   	metricDimension: &metricDimensionProperty{
//   		dimensionName: jsii.String("dimensionName"),
//
//   		// the properties below are optional
//   		operator: jsii.String("operator"),
//   	},
//   }
//
type CfnSecurityProfile_MetricToRetainProperty struct {
	// A standard of measurement.
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The dimension of the metric.
	MetricDimension interface{} `field:"optional" json:"metricDimension" yaml:"metricDimension"`
}

