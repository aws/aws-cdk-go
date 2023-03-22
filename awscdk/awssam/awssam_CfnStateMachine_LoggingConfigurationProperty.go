package awssam


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
	// `CfnStateMachine.LoggingConfigurationProperty.Destinations`.
	Destinations interface{} `field:"required" json:"destinations" yaml:"destinations"`
	// `CfnStateMachine.LoggingConfigurationProperty.IncludeExecutionData`.
	IncludeExecutionData interface{} `field:"required" json:"includeExecutionData" yaml:"includeExecutionData"`
	// `CfnStateMachine.LoggingConfigurationProperty.Level`.
	Level *string `field:"required" json:"level" yaml:"level"`
}

