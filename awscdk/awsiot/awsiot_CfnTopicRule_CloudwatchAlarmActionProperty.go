package awsiot


// Describes an action that updates a CloudWatch alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudwatchAlarmActionProperty := &cloudwatchAlarmActionProperty{
//   	alarmName: jsii.String("alarmName"),
//   	roleArn: jsii.String("roleArn"),
//   	stateReason: jsii.String("stateReason"),
//   	stateValue: jsii.String("stateValue"),
//   }
//
type CfnTopicRule_CloudwatchAlarmActionProperty struct {
	// The CloudWatch alarm name.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// The IAM role that allows access to the CloudWatch alarm.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The reason for the alarm change.
	StateReason *string `field:"required" json:"stateReason" yaml:"stateReason"`
	// The value of the alarm state.
	//
	// Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
	StateValue *string `field:"required" json:"stateValue" yaml:"stateValue"`
}

