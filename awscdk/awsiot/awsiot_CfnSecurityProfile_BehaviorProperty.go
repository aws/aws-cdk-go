package awsiot


// A Device Defender security profile behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   behaviorProperty := &behaviorProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	criteria: &behaviorCriteriaProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		consecutiveDatapointsToAlarm: jsii.Number(123),
//   		consecutiveDatapointsToClear: jsii.Number(123),
//   		durationSeconds: jsii.Number(123),
//   		mlDetectionConfig: &machineLearningDetectionConfigProperty{
//   			confidenceLevel: jsii.String("confidenceLevel"),
//   		},
//   		statisticalThreshold: &statisticalThresholdProperty{
//   			statistic: jsii.String("statistic"),
//   		},
//   		value: &metricValueProperty{
//   			cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   			count: jsii.String("count"),
//   			number: jsii.Number(123),
//   			numbers: []interface{}{
//   				jsii.Number(123),
//   			},
//   			ports: []interface{}{
//   				jsii.Number(123),
//   			},
//   			strings: []*string{
//   				jsii.String("strings"),
//   			},
//   		},
//   	},
//   	metric: jsii.String("metric"),
//   	metricDimension: &metricDimensionProperty{
//   		dimensionName: jsii.String("dimensionName"),
//
//   		// the properties below are optional
//   		operator: jsii.String("operator"),
//   	},
//   	suppressAlerts: jsii.Boolean(false),
//   }
//
type CfnSecurityProfile_BehaviorProperty struct {
	// The name you've given to the behavior.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The criteria that determine if a device is behaving normally in regard to the `metric` .
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// What is measured by the behavior.
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// The dimension of the metric.
	MetricDimension interface{} `field:"optional" json:"metricDimension" yaml:"metricDimension"`
	// The alert status.
	//
	// If you set the value to `true` , alerts will be suppressed.
	SuppressAlerts interface{} `field:"optional" json:"suppressAlerts" yaml:"suppressAlerts"`
}

