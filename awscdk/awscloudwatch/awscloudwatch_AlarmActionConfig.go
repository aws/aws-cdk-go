package awscloudwatch


// Properties for an alarm action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmActionConfig := &alarmActionConfig{
//   	alarmActionArn: jsii.String("alarmActionArn"),
//   }
//
type AlarmActionConfig struct {
	// Return the ARN that should be used for a CloudWatch Alarm action.
	AlarmActionArn *string `field:"required" json:"alarmActionArn" yaml:"alarmActionArn"`
}

