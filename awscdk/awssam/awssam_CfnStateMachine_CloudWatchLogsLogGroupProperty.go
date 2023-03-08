package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsLogGroupProperty := &cloudWatchLogsLogGroupProperty{
//   	logGroupArn: jsii.String("logGroupArn"),
//   }
//
type CfnStateMachine_CloudWatchLogsLogGroupProperty struct {
	// `CfnStateMachine.CloudWatchLogsLogGroupProperty.LogGroupArn`.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

