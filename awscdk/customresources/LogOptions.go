package customresources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Log Options for the state machine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroupRef ILogGroupRef
//
//   logOptions := &LogOptions{
//   	Destination: logGroupRef,
//   	IncludeExecutionData: jsii.Boolean(false),
//   	Level: awscdk.Aws_stepfunctions.LogLevel_OFF,
//   }
//
type LogOptions struct {
	// The log group where the execution history events will be logged.
	// Default: - a new log group will be created.
	//
	Destination interfacesawslogs.ILogGroupRef `field:"optional" json:"destination" yaml:"destination"`
	// Determines whether execution data is included in your log.
	// Default: - false.
	//
	IncludeExecutionData *bool `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Defines which category of execution history events are logged.
	// Default: - ERROR.
	//
	Level awsstepfunctions.LogLevel `field:"optional" json:"level" yaml:"level"`
}

