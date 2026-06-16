package interfacesawscloudwatch


// A reference to a LogAlarm resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logAlarmReference := &LogAlarmReference{
//   	AlarmName: jsii.String("alarmName"),
//   	LogAlarmArn: jsii.String("logAlarmArn"),
//   }
//
type LogAlarmReference struct {
	// The AlarmName of the LogAlarm resource.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// The ARN of the LogAlarm resource.
	LogAlarmArn *string `field:"required" json:"logAlarmArn" yaml:"logAlarmArn"`
}

