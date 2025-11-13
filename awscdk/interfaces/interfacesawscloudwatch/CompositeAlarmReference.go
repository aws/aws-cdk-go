package interfacesawscloudwatch


// A reference to a CompositeAlarm resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   compositeAlarmReference := &CompositeAlarmReference{
//   	AlarmName: jsii.String("alarmName"),
//   	CompositeAlarmArn: jsii.String("compositeAlarmArn"),
//   }
//
type CompositeAlarmReference struct {
	// The AlarmName of the CompositeAlarm resource.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// The ARN of the CompositeAlarm resource.
	CompositeAlarmArn *string `field:"required" json:"compositeAlarmArn" yaml:"compositeAlarmArn"`
}

