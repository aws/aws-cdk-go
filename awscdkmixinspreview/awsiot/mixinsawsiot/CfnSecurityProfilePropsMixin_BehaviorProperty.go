package mixinsawsiot


// A Device Defender security profile behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   behaviorProperty := &BehaviorProperty{
//   	Criteria: &BehaviorCriteriaProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		ConsecutiveDatapointsToAlarm: jsii.Number(123),
//   		ConsecutiveDatapointsToClear: jsii.Number(123),
//   		DurationSeconds: jsii.Number(123),
//   		MlDetectionConfig: &MachineLearningDetectionConfigProperty{
//   			ConfidenceLevel: jsii.String("confidenceLevel"),
//   		},
//   		StatisticalThreshold: &StatisticalThresholdProperty{
//   			Statistic: jsii.String("statistic"),
//   		},
//   		Value: &MetricValueProperty{
//   			Cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   			Count: jsii.String("count"),
//   			Number: jsii.Number(123),
//   			Numbers: []interface{}{
//   				jsii.Number(123),
//   			},
//   			Ports: []interface{}{
//   				jsii.Number(123),
//   			},
//   			Strings: []*string{
//   				jsii.String("strings"),
//   			},
//   		},
//   	},
//   	ExportMetric: jsii.Boolean(false),
//   	Metric: jsii.String("metric"),
//   	MetricDimension: &MetricDimensionProperty{
//   		DimensionName: jsii.String("dimensionName"),
//   		Operator: jsii.String("operator"),
//   	},
//   	Name: jsii.String("name"),
//   	SuppressAlerts: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behavior.html
//
type CfnSecurityProfilePropsMixin_BehaviorProperty struct {
	// The criteria that determine if a device is behaving normally in regard to the `metric` .
	//
	// > In the AWS IoT console, you can choose to be sent an alert through Amazon SNS when AWS IoT Device Defender detects that a device is behaving anomalously.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behavior.html#cfn-iot-securityprofile-behavior-criteria
	//
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// Value indicates exporting metrics related to the behavior when it is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behavior.html#cfn-iot-securityprofile-behavior-exportmetric
	//
	ExportMetric interface{} `field:"optional" json:"exportMetric" yaml:"exportMetric"`
	// What is measured by the behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behavior.html#cfn-iot-securityprofile-behavior-metric
	//
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// The dimension of the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behavior.html#cfn-iot-securityprofile-behavior-metricdimension
	//
	MetricDimension interface{} `field:"optional" json:"metricDimension" yaml:"metricDimension"`
	// The name you've given to the behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behavior.html#cfn-iot-securityprofile-behavior-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The alert status.
	//
	// If you set the value to `true` , alerts will be suppressed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behavior.html#cfn-iot-securityprofile-behavior-suppressalerts
	//
	SuppressAlerts interface{} `field:"optional" json:"suppressAlerts" yaml:"suppressAlerts"`
}

