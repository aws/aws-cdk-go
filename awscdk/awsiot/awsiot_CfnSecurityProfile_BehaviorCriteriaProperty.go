package awsiot


// The criteria by which the behavior is determined to be normal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   behaviorCriteriaProperty := &behaviorCriteriaProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	consecutiveDatapointsToAlarm: jsii.Number(123),
//   	consecutiveDatapointsToClear: jsii.Number(123),
//   	durationSeconds: jsii.Number(123),
//   	mlDetectionConfig: &machineLearningDetectionConfigProperty{
//   		confidenceLevel: jsii.String("confidenceLevel"),
//   	},
//   	statisticalThreshold: &statisticalThresholdProperty{
//   		statistic: jsii.String("statistic"),
//   	},
//   	value: &metricValueProperty{
//   		cidrs: []*string{
//   			jsii.String("cidrs"),
//   		},
//   		count: jsii.String("count"),
//   		number: jsii.Number(123),
//   		numbers: []interface{}{
//   			jsii.Number(123),
//   		},
//   		ports: []interface{}{
//   			jsii.Number(123),
//   		},
//   		strings: []*string{
//   			jsii.String("strings"),
//   		},
//   	},
//   }
//
type CfnSecurityProfile_BehaviorCriteriaProperty struct {
	// The operator that relates the thing measured ( `metric` ) to the criteria (containing a `value` or `statisticalThreshold` ).
	//
	// Valid operators include:
	//
	// - `string-list` : `in-set` and `not-in-set`
	// - `number-list` : `in-set` and `not-in-set`
	// - `ip-address-list` : `in-cidr-set` and `not-in-cidr-set`
	// - `number` : `less-than` , `less-than-equals` , `greater-than` , and `greater-than-equals`.
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs.
	//
	// If not specified, the default is 1.
	ConsecutiveDatapointsToAlarm *float64 `field:"optional" json:"consecutiveDatapointsToAlarm" yaml:"consecutiveDatapointsToAlarm"`
	// If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared.
	//
	// If not specified, the default is 1.
	ConsecutiveDatapointsToClear *float64 `field:"optional" json:"consecutiveDatapointsToClear" yaml:"consecutiveDatapointsToClear"`
	// Use this to specify the time duration over which the behavior is evaluated, for those criteria that have a time dimension (for example, `NUM_MESSAGES_SENT` ).
	//
	// For a `statisticalThreshhold` metric comparison, measurements from all devices are accumulated over this time duration before being used to calculate percentiles, and later, measurements from an individual device are also accumulated over this time duration before being given a percentile rank. Cannot be used with list-based metric datatypes.
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// The confidence level of the detection model.
	MlDetectionConfig interface{} `field:"optional" json:"mlDetectionConfig" yaml:"mlDetectionConfig"`
	// A statistical ranking (percentile)that indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.
	StatisticalThreshold interface{} `field:"optional" json:"statisticalThreshold" yaml:"statisticalThreshold"`
	// The value to be compared with the `metric` .
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

