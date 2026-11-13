package awscloudwatch


// A rule that queries telemetry, evaluates the result against a condition, and sets how missing data is handled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   telemetryRuleProperty := &TelemetryRuleProperty{
//   	Condition: &ConditionProperty{
//   		Comparator: jsii.String("comparator"),
//   		CriticalThreshold: jsii.Number(123),
//   		ThresholdField: jsii.String("thresholdField"),
//   		ThresholdMode: jsii.String("thresholdMode"),
//   		WarningThreshold: jsii.Number(123),
//   	},
//   	Evaluation: &EvaluationProperty{
//   		IntervalSeconds: jsii.Number(123),
//   		PendingDurationSeconds: jsii.Number(123),
//   		RecoveryDurationSeconds: jsii.Number(123),
//   	},
//   	NoData: &NoDataProperty{
//   		TreatAs: jsii.String("treatAs"),
//   	},
//   	Query: &QueryProperty{
//   		Expression: jsii.String("expression"),
//   		Language: jsii.String("language"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-telemetryrule.html
//
type CfnAlertPropsMixin_TelemetryRuleProperty struct {
	// The thresholds that set the alert's state.
	//
	// Specify at least one of WarningThreshold and CriticalThreshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-telemetryrule.html#cfn-cloudwatch-alert-telemetryrule-condition
	//
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// How often the alert is evaluated, and how long a change must last before the state changes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-telemetryrule.html#cfn-cloudwatch-alert-telemetryrule-evaluation
	//
	Evaluation interface{} `field:"optional" json:"evaluation" yaml:"evaluation"`
	// How the alert behaves when an evaluation returns no data.
	//
	// If omitted, the alert reports NODATA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-telemetryrule.html#cfn-cloudwatch-alert-telemetryrule-nodata
	//
	NoData interface{} `field:"optional" json:"noData" yaml:"noData"`
	// The query the alert runs, and the language it is written in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-telemetryrule.html#cfn-cloudwatch-alert-telemetryrule-query
	//
	Query interface{} `field:"optional" json:"query" yaml:"query"`
}

