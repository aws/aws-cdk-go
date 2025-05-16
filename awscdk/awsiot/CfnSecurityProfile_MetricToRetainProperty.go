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
//   metricToRetainProperty := &MetricToRetainProperty{
//   	Metric: jsii.String("metric"),
//
//   	// the properties below are optional
//   	ExportMetric: jsii.Boolean(false),
//   	MetricDimension: &MetricDimensionProperty{
//   		DimensionName: jsii.String("dimensionName"),
//
//   		// the properties below are optional
//   		Operator: jsii.String("operator"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metrictoretain.html
//
type CfnSecurityProfile_MetricToRetainProperty struct {
	// A standard of measurement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metrictoretain.html#cfn-iot-securityprofile-metrictoretain-metric
	//
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The value indicates exporting metrics related to the `MetricToRetain` when it's true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metrictoretain.html#cfn-iot-securityprofile-metrictoretain-exportmetric
	//
	ExportMetric interface{} `field:"optional" json:"exportMetric" yaml:"exportMetric"`
	// The dimension of the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metrictoretain.html#cfn-iot-securityprofile-metrictoretain-metricdimension
	//
	MetricDimension interface{} `field:"optional" json:"metricDimension" yaml:"metricDimension"`
}

