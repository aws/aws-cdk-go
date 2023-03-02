package awsstepfunctions


// Defines a destination for `LoggingConfiguration` .
//
// > For more information on logging with `EXPRESS` workflows, see [Logging Express Workflows Using CloudWatch Logs](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDestinationProperty := &logDestinationProperty{
//   	cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   		logGroupArn: jsii.String("logGroupArn"),
//   	},
//   }
//
type CfnStateMachine_LogDestinationProperty struct {
	// An object describing a CloudWatch log group.
	//
	// For more information, see [AWS::Logs::LogGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html) in the AWS CloudFormation User Guide.
	CloudWatchLogsLogGroup interface{} `field:"optional" json:"cloudWatchLogsLogGroup" yaml:"cloudWatchLogsLogGroup"`
}

