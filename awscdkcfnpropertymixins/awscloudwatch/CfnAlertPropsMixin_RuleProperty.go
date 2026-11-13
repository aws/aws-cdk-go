package awscloudwatch


// The rule that defines how the alert is evaluated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ruleProperty := &RuleProperty{
//   	TelemetryRule: &TelemetryRuleProperty{
//   		Condition: &ConditionProperty{
//   			Comparator: jsii.String("comparator"),
//   			CriticalThreshold: jsii.Number(123),
//   			ThresholdField: jsii.String("thresholdField"),
//   			ThresholdMode: jsii.String("thresholdMode"),
//   			WarningThreshold: jsii.Number(123),
//   		},
//   		Evaluation: &EvaluationProperty{
//   			IntervalSeconds: jsii.Number(123),
//   			PendingDurationSeconds: jsii.Number(123),
//   			RecoveryDurationSeconds: jsii.Number(123),
//   		},
//   		NoData: &NoDataProperty{
//   			TreatAs: jsii.String("treatAs"),
//   		},
//   		Query: &QueryProperty{
//   			Expression: jsii.String("expression"),
//   			Language: jsii.String("language"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-rule.html
//
type CfnAlertPropsMixin_RuleProperty struct {
	// A rule that queries telemetry, evaluates the result against a condition, and sets how missing data is handled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-rule.html#cfn-cloudwatch-alert-rule-telemetryrule
	//
	TelemetryRule interface{} `field:"optional" json:"telemetryRule" yaml:"telemetryRule"`
}

