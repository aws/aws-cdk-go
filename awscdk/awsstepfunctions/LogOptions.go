package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Defines what execution history events are logged and where they are logged.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))
//
//   definition := sfn.Chain_Start(sfn.NewPass(this, jsii.String("Pass")))
//
//   sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &StateMachineProps{
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
//   	Logs: &LogOptions{
//   		Destination: logGroup,
//   		Level: sfn.LogLevel_ALL,
//   	},
//   })
//
type LogOptions struct {
	// The log group where the execution history events will be logged.
	Destination awslogs.ILogGroup `field:"required" json:"destination" yaml:"destination"`
	// Determines whether execution data is included in your log.
	// Default: false.
	//
	IncludeExecutionData *bool `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Defines which category of execution history events are logged.
	// Default: ERROR.
	//
	Level LogLevel `field:"optional" json:"level" yaml:"level"`
}

