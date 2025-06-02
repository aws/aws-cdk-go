package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDestinationProperty := &LogDestinationProperty{
//   	CloudWatchLogsLogGroup: &CloudWatchLogsLogGroupProperty{
//   		LogGroupArn: jsii.String("logGroupArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-logdestination.html
//
type CfnStateMachine_LogDestinationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-logdestination.html#cfn-serverless-statemachine-logdestination-cloudwatchlogsloggroup
	//
	CloudWatchLogsLogGroup interface{} `field:"required" json:"cloudWatchLogsLogGroup" yaml:"cloudWatchLogsLogGroup"`
}

