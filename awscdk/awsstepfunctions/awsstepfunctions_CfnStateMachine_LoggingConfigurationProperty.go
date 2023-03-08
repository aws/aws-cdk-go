package awsstepfunctions


// Defines what execution history events are logged and where they are logged.
//
// > By default, the `level` is set to `OFF` . For more information see [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationProperty := &loggingConfigurationProperty{
//   	destinations: []interface{}{
//   		&logDestinationProperty{
//   			cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   				logGroupArn: jsii.String("logGroupArn"),
//   			},
//   		},
//   	},
//   	includeExecutionData: jsii.Boolean(false),
//   	level: jsii.String("level"),
//   }
//
type CfnStateMachine_LoggingConfigurationProperty struct {
	// An array of objects that describes where your execution history events will be logged.
	//
	// Limited to size 1. Required, if your log level is not set to `OFF` .
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Determines whether execution data is included in your log.
	//
	// When set to `false` , data is excluded.
	IncludeExecutionData interface{} `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Defines which category of execution history events are logged.
	Level *string `field:"optional" json:"level" yaml:"level"`
}

