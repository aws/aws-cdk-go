package awscloudwatch


// The thresholds that set the alert's state.
//
// Specify at least one of WarningThreshold and CriticalThreshold.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   conditionProperty := &ConditionProperty{
//   	Comparator: jsii.String("comparator"),
//   	CriticalThreshold: jsii.Number(123),
//   	ThresholdField: jsii.String("thresholdField"),
//   	ThresholdMode: jsii.String("thresholdMode"),
//   	WarningThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-condition.html
//
type CfnAlertPropsMixin_ConditionProperty struct {
	// The comparison operator.
	//
	// Valid values: GT, GTE, LT, LTE (greater than, greater than or equal to, less than, less than or equal to).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-condition.html#cfn-cloudwatch-alert-condition-comparator
	//
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// The value at which the alert enters the CRITICAL state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-condition.html#cfn-cloudwatch-alert-condition-criticalthreshold
	//
	CriticalThreshold *float64 `field:"optional" json:"criticalThreshold" yaml:"criticalThreshold"`
	// The result field to compare against the threshold.
	//
	// Required when ThresholdMode is FIELD_VALUE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-condition.html#cfn-cloudwatch-alert-condition-thresholdfield
	//
	ThresholdField *string `field:"optional" json:"thresholdField" yaml:"thresholdField"`
	// How the threshold is applied.
	//
	// COUNT_OF_RESULTS compares the number of rows the query returns. FIELD_VALUE compares the value of ThresholdField in each result row; each row is tracked separately.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-condition.html#cfn-cloudwatch-alert-condition-thresholdmode
	//
	ThresholdMode *string `field:"optional" json:"thresholdMode" yaml:"thresholdMode"`
	// The value at which the alert enters the WARNING state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-condition.html#cfn-cloudwatch-alert-condition-warningthreshold
	//
	WarningThreshold *float64 `field:"optional" json:"warningThreshold" yaml:"warningThreshold"`
}

