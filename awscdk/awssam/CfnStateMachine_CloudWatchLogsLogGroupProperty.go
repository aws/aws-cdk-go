package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsLogGroupProperty := &CloudWatchLogsLogGroupProperty{
//   	LogGroupArn: jsii.String("logGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-cloudwatchlogsloggroup.html
//
type CfnStateMachine_CloudWatchLogsLogGroupProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-cloudwatchlogsloggroup.html#cfn-serverless-statemachine-cloudwatchlogsloggroup-loggrouparn
	//
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

