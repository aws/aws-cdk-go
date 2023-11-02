package awscdkscheduleralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   var target lambdaInvoke
//
//
//   oneTimeSchedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(cdk.Duration_Hours(jsii.Number(12))),
//   	Target: Target,
//   	TargetOverrides: &ScheduleTargetProps{
//   		Input: awscdkscheduleralpha.ScheduleTargetInput_FromText(jsii.String("Overriding Target Input")),
//   		MaxEventAge: awscdk.Duration_Seconds(jsii.Number(180)),
//   		RetryAttempts: jsii.Number(5),
//   	},
//   })
//
// Experimental.
type ScheduleTargetProps struct {
	// The text, or well-formed JSON, passed to the target.
	//
	// If you are configuring a templated Lambda, AWS Step Functions, or Amazon EventBridge target,
	// the input must be a well-formed JSON. For all other target types, a JSON is not required.
	// Default: - The target's input is used.
	//
	// Experimental.
	Input ScheduleTargetInput `field:"optional" json:"input" yaml:"input"`
	// The maximum amount of time, in seconds, to continue to make retry attempts.
	// Default: - The target's maximumEventAgeInSeconds is used.
	//
	// Experimental.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of retry attempts to make before the request fails.
	// Default: - The target's maximumRetryAttempts is used.
	//
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
}

