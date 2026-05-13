package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for creating a PromQL Alarm.
//
// Example:
//   cloudwatch.NewPromQLAlarm(this, jsii.String("HighLatencyAlarm"), &PromQLAlarmProps{
//   	AlarmDescription: jsii.String("P99 latency exceeds 500ms for 5 minutes"),
//   	Query: jsii.String("histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m])) > 0.5"),
//   	EvaluationInterval: awscdk.Duration_Seconds(jsii.Number(60)),
//   	PendingPeriod: awscdk.Duration_*Seconds(jsii.Number(300)),
//   	RecoveryPeriod: awscdk.Duration_*Seconds(jsii.Number(600)),
//   })
//
type PromQLAlarmProps struct {
	// The frequency at which the alarm is evaluated.
	//
	// Must be between 10 seconds and 3600 seconds.
	EvaluationInterval awscdk.Duration `field:"required" json:"evaluationInterval" yaml:"evaluationInterval"`
	// The PromQL query that the alarm evaluates.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Whether the actions for this alarm are enabled.
	// Default: true.
	//
	ActionsEnabled *bool `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// Description for the alarm.
	// Default: - No description.
	//
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// Name of the alarm.
	// Default: - Automatically generated name.
	//
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// The duration that a contributor must continuously breach before the contributor transitions to ALARM state.
	// Default: - No pending period.
	//
	PendingPeriod awscdk.Duration `field:"optional" json:"pendingPeriod" yaml:"pendingPeriod"`
	// The duration that a contributor must continuously not be breaching before it transitions back to the OK state.
	// Default: - No recovery period.
	//
	RecoveryPeriod awscdk.Duration `field:"optional" json:"recoveryPeriod" yaml:"recoveryPeriod"`
}

