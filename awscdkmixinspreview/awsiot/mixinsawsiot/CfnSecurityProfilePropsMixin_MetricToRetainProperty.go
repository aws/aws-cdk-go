package mixinsawsiot


// The metric you want to retain.
//
// Dimensions are optional.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricToRetainProperty := &MetricToRetainProperty{
//   	ExportMetric: jsii.Boolean(false),
//   	Metric: jsii.String("metric"),
//   	MetricDimension: &MetricDimensionProperty{
//   		DimensionName: jsii.String("dimensionName"),
//   		Operator: jsii.String("operator"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metrictoretain.html
//
type CfnSecurityProfilePropsMixin_MetricToRetainProperty struct {
	// The value indicates exporting metrics related to the `MetricToRetain` when it's true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metrictoretain.html#cfn-iot-securityprofile-metrictoretain-exportmetric
	//
	ExportMetric interface{} `field:"optional" json:"exportMetric" yaml:"exportMetric"`
	// A standard of measurement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metrictoretain.html#cfn-iot-securityprofile-metrictoretain-metric
	//
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// The dimension of the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metrictoretain.html#cfn-iot-securityprofile-metrictoretain-metricdimension
	//
	MetricDimension interface{} `field:"optional" json:"metricDimension" yaml:"metricDimension"`
}

