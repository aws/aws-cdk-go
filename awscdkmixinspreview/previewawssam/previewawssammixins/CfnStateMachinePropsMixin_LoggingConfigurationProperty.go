package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnStateMachinePropsMixin_LoggingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-loggingconfiguration.html#cfn-serverless-statemachine-loggingconfiguration-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-loggingconfiguration.html#cfn-serverless-statemachine-loggingconfiguration-includeexecutiondata
	//
	IncludeExecutionData interface{} `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-loggingconfiguration.html#cfn-serverless-statemachine-loggingconfiguration-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
}

