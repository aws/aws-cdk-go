package awssam


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
	// `CfnStateMachine.LogDestinationProperty.CloudWatchLogsLogGroup`.
	CloudWatchLogsLogGroup interface{} `field:"required" json:"cloudWatchLogsLogGroup" yaml:"cloudWatchLogsLogGroup"`
}

