package awsstepfunctions


// Defines what execution history events are logged and where they are logged.
//
// Step Functions provides the log levels — `OFF` , `ALL` , `ERROR` , and `FATAL` . No event types log when set to `OFF` and all event types do when set to `ALL` .
//
// > By default, the `level` is set to `OFF` . For more information see [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationProperty := &LoggingConfigurationProperty{
//   	Destinations: []interface{}{
//   		&LogDestinationProperty{
//   			CloudWatchLogsLogGroup: &CloudWatchLogsLogGroupProperty{
//   				LogGroupArn: jsii.String("logGroupArn"),
//   			},
//   		},
//   	},
//   	IncludeExecutionData: jsii.Boolean(false),
//   	Level: jsii.String("level"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html
//
type CfnStateMachine_LoggingConfigurationProperty struct {
	// An array of objects that describes where your execution history events will be logged.
	//
	// Limited to size 1. Required, if your log level is not set to `OFF` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Determines whether execution data is included in your log.
	//
	// When set to `false` , data is excluded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-includeexecutiondata
	//
	IncludeExecutionData interface{} `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Defines which category of execution history events are logged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
}

