package customresources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Initialization properties for the `WaiterStateMachine` construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var logGroup logGroup
//
//   waiterStateMachineProps := &WaiterStateMachineProps{
//   	BackoffRate: jsii.Number(123),
//   	Interval: cdk.Duration_Minutes(jsii.Number(30)),
//   	IsCompleteHandler: function_,
//   	MaxAttempts: jsii.Number(123),
//   	TimeoutHandler: function_,
//
//   	// the properties below are optional
//   	DisableLogging: jsii.Boolean(false),
//   	LogOptions: &LogOptions{
//   		Destination: logGroup,
//   		IncludeExecutionData: jsii.Boolean(false),
//   		Level: awscdk.Aws_stepfunctions.LogLevel_OFF,
//   	},
//   }
//
type WaiterStateMachineProps struct {
	// Backoff between attempts.
	BackoffRate *float64 `field:"required" json:"backoffRate" yaml:"backoffRate"`
	// The interval to wait between attempts.
	Interval awscdk.Duration `field:"required" json:"interval" yaml:"interval"`
	// The main handler that notifies if the waiter to decide 'complete' or 'incomplete'.
	IsCompleteHandler awslambda.IFunction `field:"required" json:"isCompleteHandler" yaml:"isCompleteHandler"`
	// Number of attempts.
	MaxAttempts *float64 `field:"required" json:"maxAttempts" yaml:"maxAttempts"`
	// The handler to call if the waiter times out and is incomplete.
	TimeoutHandler awslambda.IFunction `field:"required" json:"timeoutHandler" yaml:"timeoutHandler"`
	// Whether logging for the state machine is disabled.
	// Default: - false.
	//
	DisableLogging *bool `field:"optional" json:"disableLogging" yaml:"disableLogging"`
	// Defines what execution history events are logged and where they are logged.
	// Default: - A default log group will be created if logging is enabled.
	//
	LogOptions *LogOptions `field:"optional" json:"logOptions" yaml:"logOptions"`
}

