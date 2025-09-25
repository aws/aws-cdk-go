package awslightsail


// A reference to a Alarm resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmReference := &AlarmReference{
//   	AlarmArn: jsii.String("alarmArn"),
//   	AlarmName: jsii.String("alarmName"),
//   }
//
type AlarmReference struct {
	// The ARN of the Alarm resource.
	AlarmArn *string `field:"required" json:"alarmArn" yaml:"alarmArn"`
	// The AlarmName of the Alarm resource.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
}

