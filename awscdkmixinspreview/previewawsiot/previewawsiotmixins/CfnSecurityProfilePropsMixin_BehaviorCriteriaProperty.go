package previewawsiotmixins


// The criteria by which the behavior is determined to be normal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   behaviorCriteriaProperty := &BehaviorCriteriaProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	ConsecutiveDatapointsToAlarm: jsii.Number(123),
//   	ConsecutiveDatapointsToClear: jsii.Number(123),
//   	DurationSeconds: jsii.Number(123),
//   	MlDetectionConfig: &MachineLearningDetectionConfigProperty{
//   		ConfidenceLevel: jsii.String("confidenceLevel"),
//   	},
//   	StatisticalThreshold: &StatisticalThresholdProperty{
//   		Statistic: jsii.String("statistic"),
//   	},
//   	Value: &MetricValueProperty{
//   		Cidrs: []*string{
//   			jsii.String("cidrs"),
//   		},
//   		Count: jsii.String("count"),
//   		Number: jsii.Number(123),
//   		Numbers: []interface{}{
//   			jsii.Number(123),
//   		},
//   		Ports: []interface{}{
//   			jsii.Number(123),
//   		},
//   		Strings: []*string{
//   			jsii.String("strings"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html
//
type CfnSecurityProfilePropsMixin_BehaviorCriteriaProperty struct {
	// The operator that relates the thing measured ( `metric` ) to the criteria (containing a `value` or `statisticalThreshold` ).
	//
	// Valid operators include:
	//
	// - `string-list` : `in-set` and `not-in-set`
	// - `number-list` : `in-set` and `not-in-set`
	// - `ip-address-list` : `in-cidr-set` and `not-in-cidr-set`
	// - `number` : `less-than` , `less-than-equals` , `greater-than` , and `greater-than-equals`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-comparisonoperator
	//
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs.
	//
	// If not specified, the default is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-consecutivedatapointstoalarm
	//
	ConsecutiveDatapointsToAlarm *float64 `field:"optional" json:"consecutiveDatapointsToAlarm" yaml:"consecutiveDatapointsToAlarm"`
	// If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared.
	//
	// If not specified, the default is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-consecutivedatapointstoclear
	//
	ConsecutiveDatapointsToClear *float64 `field:"optional" json:"consecutiveDatapointsToClear" yaml:"consecutiveDatapointsToClear"`
	// Use this to specify the time duration over which the behavior is evaluated, for those criteria that have a time dimension (for example, `NUM_MESSAGES_SENT` ).
	//
	// For a `statisticalThreshhold` metric comparison, measurements from all devices are accumulated over this time duration before being used to calculate percentiles, and later, measurements from an individual device are also accumulated over this time duration before being given a percentile rank. Cannot be used with list-based metric datatypes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-durationseconds
	//
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// The confidence level of the detection model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-mldetectionconfig
	//
	MlDetectionConfig interface{} `field:"optional" json:"mlDetectionConfig" yaml:"mlDetectionConfig"`
	// A statistical ranking (percentile)that indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-statisticalthreshold
	//
	StatisticalThreshold interface{} `field:"optional" json:"statisticalThreshold" yaml:"statisticalThreshold"`
	// The value to be compared with the `metric` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-behaviorcriteria.html#cfn-iot-securityprofile-behaviorcriteria-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

