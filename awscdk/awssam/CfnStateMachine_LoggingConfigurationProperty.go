package awssam


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-loggingconfiguration.html
//
type CfnStateMachine_LoggingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-loggingconfiguration.html#cfn-serverless-statemachine-loggingconfiguration-destinations
	//
	Destinations interface{} `field:"required" json:"destinations" yaml:"destinations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-loggingconfiguration.html#cfn-serverless-statemachine-loggingconfiguration-includeexecutiondata
	//
	IncludeExecutionData interface{} `field:"required" json:"includeExecutionData" yaml:"includeExecutionData"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-loggingconfiguration.html#cfn-serverless-statemachine-loggingconfiguration-level
	//
	Level *string `field:"required" json:"level" yaml:"level"`
}

