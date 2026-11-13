package awscloudwatch


// How often the alert is evaluated, and how long a change must last before the state changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   evaluationProperty := &EvaluationProperty{
//   	IntervalSeconds: jsii.Number(123),
//   	PendingDurationSeconds: jsii.Number(123),
//   	RecoveryDurationSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-evaluation.html
//
type CfnAlertPropsMixin_EvaluationProperty struct {
	// The time between evaluations, in seconds.
	//
	// Valid values: 30, 60, 120, 300, 600, 900, 1800, 3600.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-evaluation.html#cfn-cloudwatch-alert-evaluation-intervalseconds
	//
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
	// How long a threshold breach must last before the alert changes state, in seconds.
	//
	// Must be a multiple of IntervalSeconds. 0 changes state immediately.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-evaluation.html#cfn-cloudwatch-alert-evaluation-pendingdurationseconds
	//
	PendingDurationSeconds *float64 `field:"optional" json:"pendingDurationSeconds" yaml:"pendingDurationSeconds"`
	// How long the result must stay within its thresholds before the alert returns to OK, in seconds.
	//
	// Must be a multiple of IntervalSeconds. 0 returns to OK immediately.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-evaluation.html#cfn-cloudwatch-alert-evaluation-recoverydurationseconds
	//
	RecoveryDurationSeconds *float64 `field:"optional" json:"recoveryDurationSeconds" yaml:"recoveryDurationSeconds"`
}

