package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
)

// Defines what execution history events are logged and where they are logged.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))
//
//   sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   	definition: sfn.chain.start(sfn.NewPass(this, jsii.String("Pass"))),
//   	logs: &logOptions{
//   		destination: logGroup,
//   		level: sfn.logLevel_ALL,
//   	},
//   })
//
// Experimental.
type LogOptions struct {
	// The log group where the execution history events will be logged.
	// Experimental.
	Destination awslogs.ILogGroup `field:"required" json:"destination" yaml:"destination"`
	// Determines whether execution data is included in your log.
	// Experimental.
	IncludeExecutionData *bool `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Defines which category of execution history events are logged.
	// Experimental.
	Level LogLevel `field:"optional" json:"level" yaml:"level"`
}

